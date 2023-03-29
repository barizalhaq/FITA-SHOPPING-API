package usecase

import (
	"errors"
	"log"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
)

func (u *cartUsecase) AddProducts(user domain.User, input gModel.EncartInput) (*domain.Cart, error) {
	var items map[domain.Product]int = map[domain.Product]int{}

	existingCart, err := u.cartRepo.View(&user)
	if err != nil {
		log.Fatalf("[cartUsecase.AddProducts] cartRepo.View Error: %v", err)
		return nil, err
	}

	if existingCart == nil {
		_, err = u.cartRepo.Create(&user)
		if err != nil {
			log.Fatalf("[cartUsecase.AddProducts] cartRepo.Create Error: %v", err)
			return nil, err
		}
	}

	for _, itemToAdd := range input.ProductsToAdd {
		product, err := u.productRepo.GetProductByID(uint64(itemToAdd.ProductID))
		if err != nil {
			log.Fatalf("[cartUsecase.AddProducts] productRepo.GetProductByID Error: %v", err)
			return nil, err
		}

		if product == nil {
			return nil, errors.New("invalid product ID")
		}

		if itemToAdd.Qty > product.Qty {
			return nil, errors.New("invalid product qty")
		}

		items[*product] = itemToAdd.Qty
	}

	return u.cartRepo.AddProducts(&user, items)
}
