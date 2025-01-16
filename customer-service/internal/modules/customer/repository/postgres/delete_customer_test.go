package customer_repository_postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCustomer(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repomock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	customerID := uuid.New()

	t.Run("should delete customer successfully", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(deleteCustomerQuery).
			WithArgs(customerID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sqlMock.ExpectCommit()

		err := repomock.DeleteCustomer(context.Background(), customerID)
		assert.Nil(t, err)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when no rows affected", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(deleteCustomerQuery).
			WithArgs(customerID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sqlMock.ExpectRollback()

		err := repomock.DeleteCustomer(context.Background(), customerID)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Customer not found")
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("should return error when DB fails", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(deleteCustomerQuery).
			WithArgs(customerID).
			WillReturnError(errors.New("db error"))

		sqlMock.ExpectRollback()

		err := repomock.DeleteCustomer(context.Background(), customerID)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "db error")
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})
}
