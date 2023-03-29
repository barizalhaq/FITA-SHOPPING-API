package usecase

import (
	"errors"
	"log"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
)

func (u *cartUsecase) DecreaseQty(user domain.User, input gModel.DecreaseCartProductQtyInput) (*domain.Cart, error) {
	var items map[domain.Product]int = map[domain.Product]int{}

	existingCart, err := u.cartRepo.View(&user)
	if err != nil {
		log.Printf("[cartUsecase.DecreaseQty] cartRepo.View Error: %v", err)
		return nil, err
	}

	if existingCart == nil {
		return nil, errors.New("no products attached to the cart")
	}

	for _, itemToDcs := range input.ProductsToAdd {
		product, err := u.productRepo.GetProductByID(uint64(itemToDcs.ProductID))
		if err != nil {
			log.Printf("[cartUsecase.DecreaseQty] productRepo.GetProductByID Error: %v", err)
			return nil, err
		}

		if product == nil {
			return nil, errors.New("invalid product ID")
		}

		items[*product] = itemToDcs.Qty
	}

	return u.cartRepo.SubtractProducts(&user, items)
}
