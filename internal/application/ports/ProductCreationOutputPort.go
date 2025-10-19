package ports

import (
	"context"

	"products/internal/domain"
)

type ProductCreationOutputPort interface {
	Save(ctx context.Context, product *domain.Product) error
}
