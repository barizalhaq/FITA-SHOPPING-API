package repository

import (
	"errors"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

func (r *pgCartRepository) View(user *domain.User) (*domain.Cart, error) {
	var (
		userCart     domain.Cart
		cartItems    []domain.CartItem
		cartProducts []domain.Product
	)

	err := r.db.Model(user).Association("Cart").Find(&userCart)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	if userCart.ID == 0 {
		return nil, nil
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
