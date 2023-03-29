package repository

import (
	"errors"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

func (r *pgUserRepository) GetPasswordByUser(user *domain.User) (string, error) {
	var hashedPass string

	err := r.db.Raw("SELECT password FROM users WHERE id = ?", user.ID).Scan(&hashedPass).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}

		return "", err
	}

	return hashedPass, nil
}
