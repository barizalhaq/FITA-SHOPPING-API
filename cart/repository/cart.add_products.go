package repository

import (
	"errors"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

func (r *pgCartRepository) AddProducts(user *domain.User, items map[domain.Product]int) (*domain.Cart, error) {
	var (
		userCart      domain.Cart
		cartItems     []domain.CartItem
		cartProducts  []domain.Product
		itemsToCreate []map[string]interface{} = []map[string]interface{}{}
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
				itemsToCreate = append(itemsToCreate, map[string]interface{}{
					"cart_id":    userCart.ID,
					"product_id": product.ID,
					"qty":        qty,
				})
				continue
			}

			subQ := r.db.Select("qty").
				Where("cart_id = ? AND product_id = ?", userCart.ID, product.ID).
				Table("cart_items")
			err = r.db.Table("cart_items").
				Where("cart_id = ? AND product_id = ?", userCart.ID, product.ID).
				Update("qty", gorm.Expr("(?) + ?", subQ, qty)).
				Error
			if err != nil {
				return err
			}
		}

		if len(itemsToCreate) > 0 {
			err = r.db.Table("cart_items").Create(&itemsToCreate).Error
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
