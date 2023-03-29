package repository

import (
	"errors"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

func (r *pgCartRepository) SubtractProducts(user *domain.User, items map[domain.Product]int) (*domain.Cart, error) {
	var (
		userCart           domain.Cart
		cartItems          []domain.CartItem
		cartProducts       []domain.Product
		productIdsToRemove []int
	)

	err := r.db.Model(user).Association("Cart").Find(&userCart)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		for product, qty := range items {
			count := r.db.Model(userCart).Where("product_id = ?", product.ID).Association("Products").Count()
			if count == 0 {
				return errors.New("invalid product ID")
			}

			var cartItemQty int
			err = r.db.Model(&domain.CartItem{}).
				Select("qty").
				Where("cart_id = ? AND product_id = ?", userCart.ID, product.ID).
				Find(&cartItemQty).Error
			if err != nil {
				return err
			}

			if qty > cartItemQty {
				return errors.New("quantity is not valid")
			}

			if qty == cartItemQty {
				productIdsToRemove = append(productIdsToRemove, int(product.ID))
				continue
			}

			subQ := r.db.Select("qty").
				Where("cart_id = ? AND product_id = ?", userCart.ID, product.ID).
				Table("cart_items")
			err = r.db.Table("cart_items").
				Where("cart_id = ? AND product_id = ?", userCart.ID, product.ID).
				Update("qty", gorm.Expr("(?) - ?", subQ, qty)).
				Error
			if err != nil {
				return err
			}
		}

		if len(productIdsToRemove) > 0 {
			_, err = r.RemoveProducts(user, productIdsToRemove)
			if err != nil {
				return err
			}
		}

		return nil
	})
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
