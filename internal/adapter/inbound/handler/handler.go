package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mattuttis/prd-project/customer-service/internal/application"
	"github.com/mattuttis/prd-project/customer-service/internal/domain"
)

type CustomerHandler struct {
	customerService *application.CustomerService
}

func NewCustomerHandler(customerService *application.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	createCustomerRequest := CreateCustomerRequest{}
	err := c.ShouldBindJSON(&createCustomerRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newCustomer, err := h.customerService.CreateCustomer(c, createCustomerRequest.FirstName, createCustomerRequest.LastName, createCustomerRequest.Age)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	customerResponse := ToCustomerResponse(newCustomer)
	c.JSON(201, customerResponse)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id format"})
		return
	}

	customer, err := h.customerService.GetCustomer(c, parsedId)
	if errors.Is(err, domain.ErrCustomerNotFound) {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	customerResponse := ToCustomerResponse(customer)
	c.JSON(200, customerResponse)
}
