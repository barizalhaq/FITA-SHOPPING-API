package usecase

import "github.com/barizalhaq/fita_shopping_api/domain"

type promoUsecase struct {
	promoRepository domain.PromoRepositoryInterface
}

func NewPromoUsecase(promoRepo domain.PromoRepositoryInterface) *promoUsecase {
	return &promoUsecase{promoRepo}
}
