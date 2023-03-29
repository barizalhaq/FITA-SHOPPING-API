package repository

import (
	"github.com/barizalhaq/fita_shopping_api/domain"
)

func (r *pgCartRepository) Create(user *domain.User) (*domain.Cart, error) {
	cart := domain.Cart{
		UserID: user.ID,
	}

	err := r.db.Model(user).Association("Cart").Append(&cart)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}
