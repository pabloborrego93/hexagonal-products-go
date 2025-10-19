package ports

import (
	"context"

	"products/internal/domain"
)

type ProductCreationCommand struct {
	Name        string
	Description string
	Price       float64
}

type ProductCreationInputPort interface {
	Create(ctx context.Context, command ProductCreationCommand) (*domain.Product, error)
}
