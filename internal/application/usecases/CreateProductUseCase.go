package usecases

import (
	"context"
	"products/internal/application/ports"
	"products/internal/domain"
)

type productCreationUseCase struct {
	outputPort ports.ProductCreationOutputPort
}

func ProductCreationUseCase(outputPort ports.ProductCreationOutputPort) ports.ProductCreationInputPort {
	return &productCreationUseCase{outputPort: outputPort}
}

func (uc *productCreationUseCase) Create(ctx context.Context, command ports.ProductCreationCommand) (*domain.Product, error) {

	p, err := domain.CreateNewProduct(command.Name, command.Description, command.Price)
	if err != nil {
		return nil, err
	}

	if err := uc.outputPort.Save(ctx, p); err != nil {
		return nil, err
	}

	return p, nil

}

var _ ports.ProductCreationInputPort = (*productCreationUseCase)(nil)
