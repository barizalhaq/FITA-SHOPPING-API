package repository

import (
	"github.com/barizalhaq/fita_shopping_api/domain"
	"gorm.io/gorm"
)

type pgProductRepository struct {
	db *gorm.DB
}

func NewPGProductRepository(db *gorm.DB) *pgProductRepository {
	return &pgProductRepository{db}
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Product{})
}
