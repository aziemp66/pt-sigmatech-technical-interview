package customer_repository

import (
	"context"
	customer_model "customer-service/internal/modules/customer/model"
	"time"

	"github.com/google/uuid"
)

// CustomerRepository defines the set of operations for interacting with customer data at the repository level.
type CustomerRepository interface {
	// CreateCustomer creates a new customer in the repository with the given details.
	// Returns the newly created customer's ID or an error if the operation fails.
	CreateCustomer(ctx context.Context, nik string, fullName string, legalName string, birthPlace string, birthDate time.Time, salary float64, idCardPhoto string, selfiePhoto string) (id string, err error)

	// GetCustomerByID retrieves a customer by their unique ID.
	// Returns the customer details wrapped in a model.Customer struct, or an error if not found.
	GetCustomerByID(ctx context.Context, id uuid.UUID) (res customer_model.Customer, err error)

	// UpdateCustomer updates the customer's details in the repository.
	// Takes the customer's ID and the updated fields.
	UpdateCustomer(ctx context.Context, id uuid.UUID, fullName string, legalName string, birthPlace string, birthDate time.Time, salary float64, idCardPhoto string, selfiePhoto string) (err error)

	// DeleteCustomer removes a customer from the repository by their ID.
	// Returns an error if the deletion operation fails.
	DeleteCustomer(ctx context.Context, id uuid.UUID) (err error)
}
