package adapters

import (
	"context"
	"sync"

	"products/internal/application/ports"
	"products/internal/domain"
)

type InMemoryProductRepository struct {
	mu    sync.RWMutex
	store map[string]*domain.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		store: make(map[string]*domain.Product),
	}
}

func (r *InMemoryProductRepository) Save(ctx context.Context, product *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[product.ID.String()] = product
	return nil
}

var _ ports.ProductCreationOutputPort = (*InMemoryProductRepository)(nil)
