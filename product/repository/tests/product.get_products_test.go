package repository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) Test_repository_Get_Products() {
	s.T().Run("Successfully get all products", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "products" LIMIT 100`),
		).WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "sku", "name", "price", "qty"},
			).AddRow(1, "DUMMY_SKU_1", "DUMMY PRODUCT 1", 5.99, 10).
				AddRow(2, "DUMMY_SKU_2", "DUMMY PRODUCT 2", 99.99, 12).
				AddRow(3, "DUMMY_SKU_3", "DUMMY PRODUCT 3", 100.99, 90),
		)

		res, err := s.repository.GetProducts()

		assert.Nil(s.T(), err)
		assert.NotNil(s.T(), res)
		assert.Equal(s.T(), 3, len(res))
	})

	s.T().Run("Failed get all products", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "products" LIMIT 100`),
		).WillReturnError(fmt.Errorf("dummy error"))

		res, err := s.repository.GetProducts()

		assert.NotNil(s.T(), err)
		assert.Nil(s.T(), res)
		assert.Equal(s.T(), 0, len(res))
	})
}
