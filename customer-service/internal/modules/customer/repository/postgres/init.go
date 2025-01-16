package customer_repository_postgres

import (
	customer_repository "customer-service/internal/modules/customer/repository"

	"github.com/jmoiron/sqlx"
)

type customerRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) customer_repository.CustomerRepository {
	return &customerRepositoryPostgres{
		db: db,
	}
}
