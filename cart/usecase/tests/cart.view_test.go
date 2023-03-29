package usecase_test

import (
	"testing"

	"github.com/barizalhaq/fita_shopping_api/cart/usecase"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestViewCart(t *testing.T) {
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

	t.Run("Successfully get cart", func(t *testing.T) {
		t.Run("With Create Cart", func(t *testing.T) {
			cartToReturn := dummyUser.Cart
			mCart.EXPECT().View(&dummyUser).
				Return(nil, nil)

			mCart.EXPECT().Create(&dummyUser).
				Return(&cartToReturn, nil).Times(1)

			cart, err := cartUc.ViewCart(dummyUser)
			assert.Nil(t, err)
			assert.NotNil(t, cart)
			assert.Equal(t, cartToReturn.UserID, cart.UserID)
		})

		t.Run("Existing Cart", func(t *testing.T) {
			cartToReturn := dummyUser.Cart
			mCart.EXPECT().View(&dummyUser).
				Return(&cartToReturn, nil)

			cart, err := cartUc.ViewCart(dummyUser)
			assert.Nil(t, err)
			assert.NotNil(t, cart)
			assert.Equal(t, cartToReturn.UserID, cart.UserID)
		})
	})
}
