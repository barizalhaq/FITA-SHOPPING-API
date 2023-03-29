package repository

import (
	"errors"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

func (r *pgUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
