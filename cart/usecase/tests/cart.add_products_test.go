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

func TestAddProducts(t *testing.T) {
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
			SKU:   "DUMMY_SKU_1",
			Name:  "DUMMY PRODUCT 1",
			Price: 10.99,
			Qty:   15,
		},
		{
			ID:    2,
			SKU:   "DUMMY_SKU_2",
			Name:  "DUMMY PRODUCT 2",
			Price: 99.99,
			Qty:   190,
		},
	}

	t.Run("Successfully add product to a cart", func(t *testing.T) {
		t.Run("With Create Cart", func(t *testing.T) {
			cartToReturn := dummyUser.Cart
			mCart.EXPECT().View(&dummyUser).
				Return(nil, nil).
				Times(1)

			mCart.EXPECT().Create(&dummyUser).
				Return(&cartToReturn, nil).Times(1)

			gomock.InOrder(
				mProduct.EXPECT().GetProductByID(dummyProducts[0].ID).
					Return(dummyProducts[0], nil),
				mProduct.EXPECT().GetProductByID(dummyProducts[1].ID).
					Return(dummyProducts[1], nil),
			)

			productsToAdd := map[domain.Product]int{
				*dummyProducts[0]: 1,
				*dummyProducts[1]: 2,
			}

			cartToReturn.Items = []domain.CartItem{
				{
					CartID:    cartToReturn.ID,
					ProductID: dummyProducts[0].ID,
				},
				{
					CartID:    cartToReturn.ID,
					ProductID: dummyProducts[1].ID,
				},
			}
			cartToReturn.Products = []domain.Product{
				*dummyProducts[0], *dummyProducts[1],
			}

			mCart.EXPECT().AddProducts(&dummyUser, productsToAdd).
				Return(&cartToReturn, nil)

			cart, err := cartUc.AddProducts(dummyUser, gModel.EncartInput{
				ProductsToAdd: []*gModel.ProductWithQty{
					{
						ProductID: int(dummyProducts[0].ID),
						Qty:       1,
					},
					{
						ProductID: int(dummyProducts[1].ID),
						Qty:       2,
					},
				},
			})
			assert.Nil(t, err)
			assert.NotNil(t, cart)
		})

		t.Run("Existing Cart", func(t *testing.T) {
			cartToReturn := dummyUser.Cart
			mCart.EXPECT().View(&dummyUser).
				Return(&cartToReturn, nil).
				Times(1)

			gomock.InOrder(
				mProduct.EXPECT().GetProductByID(dummyProducts[0].ID).
					Return(dummyProducts[0], nil),
				mProduct.EXPECT().GetProductByID(dummyProducts[1].ID).
					Return(dummyProducts[1], nil),
			)

			productsToAdd := map[domain.Product]int{
				*dummyProducts[0]: 1,
				*dummyProducts[1]: 2,
			}

			cartToReturn.Items = []domain.CartItem{
				{
					CartID:    cartToReturn.ID,
					ProductID: dummyProducts[0].ID,
				},
				{
					CartID:    cartToReturn.ID,
					ProductID: dummyProducts[1].ID,
				},
			}
			cartToReturn.Products = []domain.Product{
				*dummyProducts[0], *dummyProducts[1],
			}

			mCart.EXPECT().AddProducts(&dummyUser, productsToAdd).
				Return(&cartToReturn, nil)

			cart, err := cartUc.AddProducts(dummyUser, gModel.EncartInput{
				ProductsToAdd: []*gModel.ProductWithQty{
					{
						ProductID: int(dummyProducts[0].ID),
						Qty:       1,
					},
					{
						ProductID: int(dummyProducts[1].ID),
						Qty:       2,
					},
				},
			})
			assert.Nil(t, err)
			assert.NotNil(t, cart)
		})
	})
}
