package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/mattuttis/customer-service/internal/domain"
)

type CustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, firstName, lastName string, age int) (*domain.Customer, error) {
	customer, err := domain.NewCustomer(
		firstName,
		lastName,
		age,
	)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	return s.repo.FindByID(ctx, id)
}
