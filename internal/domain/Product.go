package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateNewProduct(name string, description string, price float64) (*Product, error) {

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if description == "" {
		return nil, errors.New("name cannot be empty")
	}

	if price <= 0 {
		return nil, errors.New("price must be positive")
	}

	now := time.Now()

	return &Product{
		ID:        uuid.New(),
		Name:      name,
		Price:     price,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil

}
