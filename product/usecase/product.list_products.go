package usecase

import "github.com/barizalhaq/fita_shopping_api/domain"

func (u *productUsecase) ListProducts() ([]domain.Product, error) {
	return u.productRepository.GetProducts()
}
