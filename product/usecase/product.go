package usecase

import "github.com/barizalhaq/fita_shopping_api/domain"

type productUsecase struct {
	productRepository domain.ProductRepositoryInterface
}

func NewProductUsecase(productRepo domain.ProductRepositoryInterface) *productUsecase {
	return &productUsecase{productRepo}
}
