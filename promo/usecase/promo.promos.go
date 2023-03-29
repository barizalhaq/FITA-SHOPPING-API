package usecase

import "github.com/barizalhaq/fita_shopping_api/domain"

func (u promoUsecase) Promos() ([]domain.Promo, error) {
	return u.promoRepository.GetPromos()
}
