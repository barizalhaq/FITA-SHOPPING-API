package repository

import (
	"fmt"

	"github.com/barizalhaq/fita_shopping_api/domain"
)

func (r *pgProductRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.Limit(100).Find(&products).Error
	fmt.Println("LEN", len(products), products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
