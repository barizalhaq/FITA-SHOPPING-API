package usecase_test

import (
	"testing"

	"github.com/barizalhaq/fita_shopping_api/cart/usecase"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDecreaseQty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mCart := mocks.NewMockCartRepositoryInterface(ctrl)
	mProduct := mocks.NewMockProductRepositoryInterface(ctrl)
	mPromo := mocks.NewMockPromoRepositoryInterface(ctrl)

	cartUc := usecase.NewCartUsecase(mCart, mProduct, mPromo)

	dummyUser := domain.User{
		ID:       1,
		Username: "DUMMY_USERNAME",
		Cart: domain.Cart{
			ID: 1,
		},
	}

	dummyProducts := []*domain.Product{
		&domain.Product{
			ID:    1,
			SKU:   "DUMMY_SKU_1",
			Name:  "DUMMY PRODUCT 1",
			Price: 10.99,
			Qty:   15,
		},
		&domain.Product{
			ID:    2,
			SKU:   "DUMMY_SKU_2",
			Name:  "DUMMY PRODUCT 2",
			Price: 99.99,
			Qty:   190,
		},
	}

	t.Run("Successfully decreasing cart items qty", func(t *testing.T) {
		cartToReturn := dummyUser.Cart
		cartToReturn.Items = []domain.CartItem{
			domain.CartItem{
				CartID:    cartToReturn.ID,
				ProductID: dummyProducts[0].ID,
				Qty:       5,
			},
			domain.CartItem{
				CartID:    cartToReturn.ID,
				ProductID: dummyProducts[1].ID,
				Qty:       8,
			},
		}
		cartToReturn.Products = []domain.Product{
			*dummyProducts[0], *dummyProducts[1],
		}

		mCart.EXPECT().View(&dummyUser).
			Return(&cartToReturn, nil)

		gomock.InOrder(
			mProduct.EXPECT().GetProductByID(dummyProducts[0].ID).
				Return(dummyProducts[0], nil),
			mProduct.EXPECT().GetProductByID(dummyProducts[1].ID).
				Return(dummyProducts[1], nil),
		)

		toSubtracts := map[domain.Product]int{
			*dummyProducts[0]: 3,
			*dummyProducts[1]: 5,
		}

		decreasedItemCart := cartToReturn
		decreasedItemCart.Items[0].Qty -= 3
		decreasedItemCart.Items[1].Qty -= 5

		mCart.EXPECT().SubtractProducts(&dummyUser, toSubtracts).
			Return(&decreasedItemCart, nil)

		cart, err := cartUc.DecreaseQty(dummyUser, gModel.DecreaseCartProductQtyInput{
			ProductsToAdd: []*gModel.ProductWithQty{
				&gModel.ProductWithQty{
					ProductID: int(dummyProducts[0].ID),
					Qty:       3,
				},
				&gModel.ProductWithQty{
					ProductID: int(dummyProducts[1].ID),
					Qty:       5,
				},
			},
		})
		assert.Nil(t, err)
		assert.NotNil(t, cart)
		assert.Equal(t, len(cartToReturn.Items), len(cart.Items))
		assert.Equal(t, len(cartToReturn.Products), len(cart.Products))
		assert.Equal(t, 2, cart.Items[0].Qty)
		assert.Equal(t, 3, cart.Items[1].Qty)
	})
}
