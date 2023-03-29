package usecase

import (
	"errors"
	"log"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
)

func (u *cartUsecase) Decart(user domain.User, input gModel.DecartInput) (*domain.Cart, error) {
	existingCart, err := u.cartRepo.View(&user)
	if err != nil {
		log.Printf("[cartUsecase.Decart] cartRepo.View Error: %v", err)
		return nil, err
	}

	if existingCart == nil {
		return nil, errors.New("no products attached to the cart")
	}

	if len(existingCart.Products) == 0 {
		return nil, errors.New("the cart is empty")
	}

	productIDsMap := map[int]struct{}{}
	for _, cartProduct := range existingCart.Products {
		productIDsMap[int(cartProduct.ID)] = struct{}{}
	}

	for _, id := range input.ProductIDs {
		if _, ok := productIDsMap[id]; !ok {
			return nil, errors.New("invalid product ID")
		}
	}

	return u.cartRepo.RemoveProducts(&user, input.ProductIDs)
}
