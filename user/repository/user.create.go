package repository

import "github.com/barizalhaq/fita_shopping_api/domain"

func (r *pgUserRepository) Create(username, password string) (*domain.User, error) {
	var user domain.User

	err := r.db.Raw(`
		INSERT INTO users (username, password)
		VALUES (?,?) RETURNING id,username
		`, username, password).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
