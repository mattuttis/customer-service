package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound error = errors.New("customer not found")
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *Customer) error
	FindByID(ctx context.Context, id uuid.UUID) (*Customer, error)
}
