package customer_repository_postgres

import (
	"context"
	util_error "customer-service/util/error"
	util_logger "customer-service/util/logger"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *customerRepositoryPostgres) UpdateCustomer(ctx context.Context, id uuid.UUID, fullName string, legalName string, birthPlace string, birthDate time.Time, salary float64, idCardPhoto string, selfiePhoto string) error {
	// Start a transaction
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	// Ensure rollback in case of panic or error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			util_logger.Fatal(ctx, fmt.Sprintf("%v", p))
		}
	}()

	// Execute the update query
	result, err := tx.ExecContext(ctx, updateCustomerQuery, id, fullName, legalName, birthPlace, birthDate, salary, idCardPhoto, selfiePhoto)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return util_error.NewNotFound(nil, "Customer not found")
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
