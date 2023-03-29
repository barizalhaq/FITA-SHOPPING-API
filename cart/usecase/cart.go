package usecase

import "github.com/barizalhaq/fita_shopping_api/domain"

type cartUsecase struct {
	cartRepo    domain.CartRepositoryInterface
	productRepo domain.ProductRepositoryInterface
	promoRepo   domain.PromoRepositoryInterface
}

func NewCartUsecase(cartRepo domain.CartRepositoryInterface, productRepo domain.ProductRepositoryInterface, promoRepo domain.PromoRepositoryInterface) *cartUsecase {
	return &cartUsecase{cartRepo, productRepo, promoRepo}
}
