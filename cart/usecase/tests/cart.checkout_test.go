package usecase_test

import (
	"reflect"
	"testing"

	"github.com/barizalhaq/fita_shopping_api/cart/usecase"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
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
		{
			ID:    1,
			SKU:   "120P90",
			Name:  "Google Home",
			Price: 49.99,
			Qty:   10,
		},
		{
			ID:    2,
			SKU:   "43N23P",
			Name:  "MacBook Pro",
			Price: 5399.99,
			Qty:   5,
		},
		{
			ID:    3,
			SKU:   "A304SD",
			Name:  "Alexa Speaker",
			Price: 109.50,
			Qty:   10,
		},
		{
			ID:    4,
			SKU:   "234234",
			Name:  "Raspberry Pi B",
			Price: 30,
			Qty:   2,
		},
	}

	t.Run("Checkout successfully", func(t *testing.T) {
		cartToReturn := dummyUser.Cart
		cartToReturn.Items = []domain.CartItem{
			{
				CartID:    cartToReturn.ID,
				ProductID: 2,
				Qty:       1,
			},
			{
				CartID:    cartToReturn.ID,
				ProductID: 4,
				Qty:       2,
			},
			{
				CartID:    cartToReturn.ID,
				ProductID: 1,
				Qty:       3,
			},
			{
				CartID:    cartToReturn.ID,
				ProductID: 3,
				Qty:       3,
			},
		}
		cartToReturn.Products = []domain.Product{
			*dummyProducts[1], *dummyProducts[3], *dummyProducts[0], *dummyProducts[2],
		}

		mCart.EXPECT().View(&dummyUser).
			Return(&cartToReturn, nil)

		promos := []domain.Promo{
			{
				SKU:            "43N23P",
				OnFree:         true,
				OnDiscount:     false,
				PurchaseAmount: 1,
				FreeProduct: struct {
					SKU string "json:\"sku\""
				}{
					SKU: "234234",
				},
			},
			{
				SKU:            "120P90",
				OnFree:         true,
				OnDiscount:     false,
				PurchaseAmount: 3,
				FreeProduct: struct {
					SKU string "json:\"sku\""
				}{
					SKU: "120P90",
				},
			},
			{
				SKU:                "A304SD",
				OnFree:             false,
				OnDiscount:         true,
				PurchaseAmount:     3,
				DiscountPercentage: 10,
			},
		}
		mPromo.EXPECT().GetPromos().
			Return(promos, nil)

		invoice, err := cartUc.Checkout(dummyUser)
		assert.Nil(t, err)
		assert.NotNil(t, invoice)
		assert.True(t, reflect.DeepEqual(invoice.Cart, cartToReturn))
		assert.Equal(t, 5825.62, invoice.TotalPrice.ActualTotalPrice)
	})
}
