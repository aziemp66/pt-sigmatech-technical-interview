package customer_repository_postgres

import (
	"context"
	util_error "customer-service/util/error"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repomock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	reqNik := "1234567890123456"
	reqFullName := "Joko Anwar"
	reqLegalName := "Joko Anwar Abadi"
	reqBirthPlace := "Aceh"
	reqBirthDate := time.Date(1945, 8, 17, 17, 17, 17, 17, time.UTC)
	reqSalary := float64(120_000_000)
	reqIdCardPhoto := "photo_id"
	reqSelfiePhoto := "photo_selfie"

	returnedID := uuid.New()

	t.Run("should insert and return id", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectQuery(getCustomerByNikQuery).
			WithArgs(reqNik).
			WillReturnError(sql.ErrNoRows)

		sqlMock.ExpectQuery(createCustomerQuery).
			WithArgs(reqNik, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto).
			WillReturnRows(sqlmock.NewRows([]string{"customer_id"}).AddRow(returnedID))

		sqlMock.ExpectCommit()

		id, err := repomock.CreateCustomer(context.Background(), reqNik, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.Nil(t, err)
		assert.Equal(t, returnedID.String(), id)

		// Verify that all expectations were met
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when customer already exists", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectQuery(getCustomerByNikQuery).
			WithArgs(reqNik).
			WillReturnRows(sqlmock.NewRows([]string{"customer_id"}).AddRow(uuid.New()))

		sqlMock.ExpectRollback()

		id, err := repomock.CreateCustomer(context.Background(), reqNik, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.NotNil(t, err)
		assert.IsType(t, util_error.NewBadRequest(nil, ""), err)
		assert.Empty(t, id)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when DB fails during NIK lookup", func(t *testing.T) {
		sqlMock.ExpectBegin()

		expectedErr := errors.New("db error")
		sqlMock.ExpectQuery(getCustomerByNikQuery).
			WithArgs(reqNik).
			WillReturnError(expectedErr)

		sqlMock.ExpectRollback()

		id, err := repomock.CreateCustomer(context.Background(), reqNik, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error())
		assert.Empty(t, id)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when DB fails during insert", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectQuery(getCustomerByNikQuery).
			WithArgs(reqNik).
			WillReturnError(sql.ErrNoRows)

		expectedErr := errors.New("insert error")
		sqlMock.ExpectQuery(createCustomerQuery).
			WithArgs(reqNik, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto).
			WillReturnError(expectedErr)

		sqlMock.ExpectRollback()

		id, err := repomock.CreateCustomer(context.Background(), reqNik, reqFullName, reqLegalName, reqBirthPlace, reqBirthDate, reqSalary, reqIdCardPhoto, reqSelfiePhoto)
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error())
		assert.Empty(t, id)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})
}
