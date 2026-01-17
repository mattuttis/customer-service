package persistence

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/mattuttis/prd-project/customer-service/internal/domain"
)

type InMemoryCustomerRepository struct {
	mu        sync.RWMutex
	customers map[uuid.UUID]*domain.Customer
}

func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[uuid.UUID]*domain.Customer),
	}
}

func (r *InMemoryCustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.customers[customer.ID] = customer
	return nil
}

func (r *InMemoryCustomerRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, err := r.customers[id]
	if !err {
		return nil, domain.ErrCustomerNotFound
	}
	return customer, nil
}
