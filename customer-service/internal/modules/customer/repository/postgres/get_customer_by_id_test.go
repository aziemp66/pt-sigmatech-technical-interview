package customer_repository_postgres

import (
	"context"
	customer_model "customer-service/internal/modules/customer/model"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerByID(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repomock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	customerID := uuid.New()
	expectedCustomer := customer_model.Customer{
		CustomerID:  customerID,
		NIK:         "1234567890123456",
		FullName:    "Joko Anwar",
		LegalName:   "Joko Anwar Abadi",
		BirthPlace:  "Aceh",
		BirthDate:   time.Date(1945, 8, 17, 17, 17, 17, 17, time.UTC),
		Salary:      120_000_000,
		IDCardPhoto: "photo_id",
		SelfiePhoto: "photo_selfie",
	}

	t.Run("should return customer by ID", func(t *testing.T) {
		sqlMock.ExpectQuery(getCustomerByIDQuery).
			WithArgs(customerID).
			WillReturnRows(sqlmock.NewRows([]string{
				"customer_id", "nik", "full_name", "legal_name", "birth_place", "birth_date", "salary", "id_card_photo", "selfie_photo",
			}).AddRow(
				expectedCustomer.CustomerID, expectedCustomer.NIK, expectedCustomer.FullName,
				expectedCustomer.LegalName, expectedCustomer.BirthPlace, expectedCustomer.BirthDate,
				expectedCustomer.Salary, expectedCustomer.IDCardPhoto, expectedCustomer.SelfiePhoto,
			))

		result, err := repomock.GetCustomerByID(context.Background(), customerID)
		assert.Nil(t, err)
		assert.Equal(t, expectedCustomer, result)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when customer not found", func(t *testing.T) {
		sqlMock.ExpectQuery(getCustomerByIDQuery).
			WithArgs(customerID).
			WillReturnError(sql.ErrNoRows)

		result, err := repomock.GetCustomerByID(context.Background(), customerID)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Customer not found")
		assert.Empty(t, result)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when DB fails", func(t *testing.T) {
		sqlMock.ExpectQuery(getCustomerByIDQuery).
			WithArgs(customerID).
			WillReturnError(errors.New("db error"))

		result, err := repomock.GetCustomerByID(context.Background(), customerID)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "db error")
		assert.Empty(t, result)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})
}
