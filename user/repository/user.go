package repository

import (
	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPGUserRepository(db *gorm.DB) *pgUserRepository {
	return &pgUserRepository{db}
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{})
}
