package domain

import (
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
)

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Username string
	Cart     Cart `gorm:"foreignKey:user_id"`
}

type UserUsecaseInterface interface {
	Register(input gModel.RegisterInput) (*User, error)
	Authenticate(input gModel.AuthenticateInput) (string, error)
}

type UserRepositoryInterface interface {
	GetUserByUsername(username string) (*User, error)
	GetPasswordByUser(user *User) (string, error)
	Create(username, password string) (*User, error)
	GetUserByID(id uint64) (*User, error)
}
