package usecase

import (
	"log"

	"github.com/barizalhaq/fita_shopping_api/domain"
)

func (u *cartUsecase) ViewCart(user domain.User) (*domain.Cart, error) {
	existingCart, err := u.cartRepo.View(&user)
	if err != nil {
		log.Fatalf("[cartUsecase.ViewCart] cartRepo.View Error: %v", err)
		return nil, err
	}

	if existingCart == nil {
		existingCart, err = u.cartRepo.Create(&user)
		if err != nil {
			log.Fatalf("[cartUsecase.ViewCart] cartRepo.Create Error: %v", err)
			return nil, err
		}
	}

	return existingCart, nil
}
