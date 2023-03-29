package usecase_test

import (
	"testing"

	"github.com/barizalhaq/fita_shopping_api/cart/usecase"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDecart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mCart := mocks.NewMockCartRepositoryInterface(ctrl)
	mProduct := mocks.NewMockProductRepositoryInterface(ctrl)
	mPromo := mocks.NewMockPromoRepositoryInterface(ctrl)

	cartUc := usecase.NewCartUsecase(mCart, mProduct, mPromo)

	dummyUser := domain.User{
		ID:       1,
		Username: "DUMMY_USERNAME",
	}

	t.Run("Successfully removing product from cart", func(t *testing.T) {
		cartToReturn := &domain.Cart{
			ID:     1,
			UserID: 1,
			Products: []domain.Product{
				domain.Product{
					ID:    1,
					SKU:   "DUMMY_SKU",
					Name:  "DUMMY PRODUCT",
					Price: 5.99,
					Qty:   5,
				},
			},
		}

		mCart.EXPECT().View(&dummyUser).
			Return(cartToReturn, nil)

		clearedCart := *cartToReturn
		clearedCart.Products = nil
		mCart.EXPECT().RemoveProducts(&dummyUser, []int{1}).
			Return(&clearedCart, nil)

		cart, err := cartUc.Decart(dummyUser, gModel.DecartInput{
			ProductIDs: []int{1},
		})
		assert.Nil(t, err)
		assert.NotNil(t, cart)
		assert.Empty(t, cart.Items)
		assert.Empty(t, cart.Products)
	})

	t.Run("Failed removing product from cart", func(t *testing.T) {
		t.Run("No Cart Exist", func(t *testing.T) {
			mCart.EXPECT().View(&dummyUser).
				Return(nil, gorm.ErrRecordNotFound)

			cart, err := cartUc.Decart(dummyUser, gModel.DecartInput{
				ProductIDs: []int{1},
			})
			assert.NotNil(t, err)
			assert.Nil(t, cart)
		})

		t.Run("No Cart Exist", func(t *testing.T) {
			cartToReturn := &domain.Cart{
				ID:       1,
				UserID:   1,
				Products: nil,
			}

			mCart.EXPECT().View(&dummyUser).
				Return(cartToReturn, nil)

			cart, err := cartUc.Decart(dummyUser, gModel.DecartInput{
				ProductIDs: []int{1},
			})
			assert.NotNil(t, err)
			assert.Nil(t, cart)
			assert.Equal(t, "the cart is empty", err.Error())
		})

		t.Run("Invalid Product ID", func(t *testing.T) {
			cartToReturn := &domain.Cart{
				ID:     1,
				UserID: 1,
				Products: []domain.Product{
					domain.Product{
						ID:    1,
						SKU:   "DUMMY_SKU",
						Name:  "DUMMY PRODUCT",
						Price: 5.99,
						Qty:   5,
					},
				},
			}

			mCart.EXPECT().View(&dummyUser).
				Return(cartToReturn, nil)

			cart, err := cartUc.Decart(dummyUser, gModel.DecartInput{
				ProductIDs: []int{5},
			})
			assert.NotNil(t, err)
			assert.Nil(t, cart)
			assert.Equal(t, "invalid product ID", err.Error())
		})
	})
}
