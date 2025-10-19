package main

import (
	"log"

	"products/internal/application/usecases"
	"products/internal/infrastructure/adapters"
)

func main() {

	outputport := adapters.NewInMemoryProductRepository()

	createUC := usecases.ProductCreationUseCase(outputport)

	handler := adapters.NewProductHandler(createUC)

	router := adapters.NewRouter(handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
