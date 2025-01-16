package customer_repository_postgres

import (
	"context"
	customer_model "customer-service/internal/modules/customer/model"
	util_error "customer-service/util/error"
	util_logger "customer-service/util/logger"
	"database/sql"
	"fmt"
	"time"
)

func (r *customerRepositoryPostgres) CreateCustomer(ctx context.Context, nik string, fullName string, legalName string, birthPlace string, birthDate time.Time, salary float64, idCardPhoto string, selfiePhoto string) (string, error) {
	// Start a transaction
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return "", err
	}

	// Ensure the transaction is rolled back in case of panic or error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			util_logger.Fatal(ctx, fmt.Sprintf("%v", p))
		}
	}()

	// Check if customer already exists
	var customer customer_model.Customer
	err = tx.GetContext(ctx, &customer, getCustomerByNikQuery, nik)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return "", err
	}
	if err == nil {
		tx.Rollback()
		return "", util_error.NewBadRequest(nil, "Nik is already used")
	}

	// Create new customer
	var id string
	err = tx.QueryRowContext(ctx, createCustomerQuery, nik, fullName, legalName, birthPlace, birthDate, salary, idCardPhoto, selfiePhoto).Scan(&id)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return "", err
	}

	return id, nil
}
