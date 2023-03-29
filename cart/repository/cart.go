package repository

import (
	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

type pgCartRepository struct {
	db *gorm.DB
}

func NewPGCartRepository(db *gorm.DB) *pgCartRepository {
	return &pgCartRepository{db}
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Cart{})
}
