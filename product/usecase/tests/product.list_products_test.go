package usecase_test

import (
	"fmt"
	"testing"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/barizalhaq/fita_shopping_api/product/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductRepositoryInterface(ctrl)

	productUC := usecase.NewProductUsecase(m)

	t.Run("Successfully get all products", func(t *testing.T) {
		dummyProducts := []domain.Product{
			domain.Product{
				ID:    1,
				SKU:   "DUMMY_SKU_1",
				Name:  "DUMMY PRODUCT 1",
				Price: 10.99,
				Qty:   15,
			},
			domain.Product{
				ID:    2,
				SKU:   "DUMMY_SKU_2",
				Name:  "DUMMY PRODUCT 2",
				Price: 99.99,
				Qty:   190,
			},
		}
		m.EXPECT().GetProducts().Return(dummyProducts, nil).Times(1)

		res, err := productUC.ListProducts()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 2, len(res))
	})

	t.Run("Successfully get empty products", func(t *testing.T) {
		m.EXPECT().GetProducts().Return([]domain.Product{}, nil).Times(1)

		res, err := productUC.ListProducts()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, len(res))
	})

	t.Run("Failed get products", func(t *testing.T) {
		m.EXPECT().GetProducts().Return(nil, fmt.Errorf("an error")).Times(1)

		res, err := productUC.ListProducts()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "an error", err.Error())
	})
}
