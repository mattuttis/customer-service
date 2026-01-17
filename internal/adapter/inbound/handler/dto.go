package handler

import "github.com/mattuttis/customer-service/internal/domain"

type CreateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required,min=1,max=100"`
	LastName  string `json:"last_name" binding:"required,min=1,max=100"`
	Age       int    `json:"age" binding:"required,gte=18,lte=150"`
}

type CustomerResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func ToCustomerResponse(c *domain.Customer) CustomerResponse {
	return CustomerResponse{
		ID:        c.ID.String(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Age:       c.Age,
	}
}
