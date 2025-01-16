package customer_repository_postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCustomer(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repomock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	customerID := uuid.New()
	reqFullName := "Updated Name"
	reqLegalName := "Updated Legal Name"
	reqBirthPlace := "Updated Place"
	reqBirthDate := time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC)
	reqSalary := float64(150_000_000)
	reqIdCardPhoto := "updated_id_photo"
	reqSelfiePhoto := "updated_selfie_photo"

	t.Run("should update customer successfully", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(updateCustomerQuery).
			WithArgs(customerID, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sqlMock.ExpectCommit()

		err := repomock.UpdateCustomer(context.Background(), customerID, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.Nil(t, err)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when no rows affected", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(updateCustomerQuery).
			WithArgs(customerID, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sqlMock.ExpectRollback()

		err := repomock.UpdateCustomer(context.Background(), customerID, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Customer not found")
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when DB fails", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(updateCustomerQuery).
			WithArgs(customerID, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto).
			WillReturnError(errors.New("db error"))

		sqlMock.ExpectRollback()

		err := repomock.UpdateCustomer(context.Background(), customerID, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "db error")
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})
}
