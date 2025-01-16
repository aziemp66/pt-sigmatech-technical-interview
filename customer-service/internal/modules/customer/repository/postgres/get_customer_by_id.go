package customer_repository_postgres

import (
	"context"
	customer_model "customer-service/internal/modules/customer/model"
	util_error "customer-service/util/error"
	"database/sql"

	"github.com/google/uuid"
)

func (r *customerRepositoryPostgres) GetCustomerByID(ctx context.Context, id uuid.UUID) (customer_model.Customer, error) {
	var customer customer_model.Customer

	// Execute query to fetch customer by ID
	err := r.db.GetContext(ctx, &customer, getCustomerByIDQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer_model.Customer{}, util_error.NewNotFound(nil, "Customer not found")
		}
		return customer_model.Customer{}, err
	}

	return customer, nil
}
