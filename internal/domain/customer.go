package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidFirstName error = errors.New("invalid first name")
	ErrInvalidLastName  error = errors.New("invalid last name")
	ErrInvalidAge       error = errors.New("invalid age, should be 18 or older")
)

type Customer struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int
}

func NewCustomer(firstName, lastName string, age int) (*Customer, error) {

	if firstName == "" {
		return nil, ErrInvalidFirstName
	}
	if lastName == "" {
		return nil, ErrInvalidLastName
	}
	if age < 18 {
		return nil, ErrInvalidAge
	}

	id := uuid.New()
	return &Customer{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}, nil
}
