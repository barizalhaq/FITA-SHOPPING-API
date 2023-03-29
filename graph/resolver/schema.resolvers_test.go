package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/domain/mocks"
	"github.com/barizalhaq/fita_shopping_api/graph"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/barizalhaq/fita_shopping_api/graph/resolver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMutationResolver_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	mUserUC := mocks.NewMockUserUsecaseInterface(ctrl)

	resolvers := resolver.Resolver{
		UserUsecase: mUserUC,
	}

	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers})))

	t.Run("Successfully registered", func(t *testing.T) {
		returnedUser := domain.User{
			ID:       1,
			Username: "dummy_user",
		}

		mUserUC.EXPECT().Register(gModel.RegisterInput{
			Username: "dummy_user",
			Password: "dummy_password",
		}).Return(&returnedUser, nil)

		var resp struct {
			Register struct {
				User struct {
					ID       int
					Username string
				}
			}
		}
		q := `
			mutation {
				register(
					input: {
						username: "dummy_user"
						password: "dummy_password"
					}
				) {
					user {
						ID
						username
					}
				}
			}
		`

		c.MustPost(q, &resp)
		assert.Equal(t, returnedUser.Username, resp.Register.User.Username)
		assert.Equal(t, int(returnedUser.ID), resp.Register.User.ID)
	})
}

func TestMutationResolver_Encart(t *testing.T) {
	ctrl := gomock.NewController(t)
	mCart := mocks.NewMockCartUsecaseInterface(ctrl)

	resolvers := resolver.Resolver{
		CartUsecase: mCart,
	}

	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers})))

	t.Run("Successfully enCart Products", func(t *testing.T) {
		products := []domain.Product{
			{
				ID:    1,
				SKU:   "DUMMY_SKU_1",
				Name:  "DUMMY PRODUCT",
				Price: 5.99,
				Qty:   5,
			},
			{
				ID:    2,
				SKU:   "DUMMY_SKU_2",
				Name:  "ANOTHER DUMMY PRODUCT",
				Price: 120.99,
				Qty:   50,
			},
		}
		cart := domain.Cart{
			ID:       1,
			UserID:   1,
			Products: products,
			Items: []domain.CartItem{
				{
					CartID:    1,
					ProductID: 1,
					Qty:       2,
				},
				{
					CartID:    1,
					ProductID: 2,
					Qty:       5,
				},
			},
		}

		user := domain.User{
			ID:       1,
			Username: "test",
		}
		mCart.EXPECT().AddProducts(user, gModel.EncartInput{
			ProductsToAdd: []*gModel.ProductWithQty{
				{
					ProductID: 1,
					Qty:       2,
				},
				{
					ProductID: 2,
					Qty:       5,
				},
			},
		}).Return(&cart, nil)

		var resp struct {
			EnCart struct {
				ID    int
				Owner struct {
					ID       int
					Username string
				}
				Products []struct {
					SKU   string
					Name  string
					Price struct {
						originalPrice float64
						currency      string
					}
					Qty int
				}
			}
		}

		q := `
			mutation {
				enCart(
					input: {
						productsToAdd: [
							{
								productID: 1
								qty: 2
							}
							{
								productID: 2
								qty: 5
							}
						]
					}
				) {
					ID
					owner {
						ID
						username
					}
					products {
						sku
						name
						price {
							originalPrice
							currency
						}
						qty
					}
				}
			}
		`

		c.MustPost(q, &resp)
	})
}
