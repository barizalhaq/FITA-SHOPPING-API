package repository

import (
	"errors"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

func (r *pgCartRepository) RemoveProducts(user *domain.User, productIds []int) (*domain.Cart, error) {
	var (
		userCart     domain.Cart
		cartProducts []domain.Product
		products     []domain.Product
		cartItems    []domain.CartItem
	)

	err := r.db.Model(user).Association("Cart").Find(&userCart)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	err = r.db.Find(&products, productIds).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&userCart).Association("Products").Delete(products)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("cart_id = ?", userCart.ID).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&userCart).Association("Products").Find(&cartProducts)
	if err != nil {
		return nil, err
	}

	userCart.Items = cartItems
	userCart.Products = cartProducts

	return &userCart, nil
}
