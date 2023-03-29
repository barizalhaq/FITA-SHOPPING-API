package resolver

import (
	"context"

	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/barizalhaq/fita_shopping_api/user/delivery/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUsecase    domain.UserUsecaseInterface
	ProductUsecase domain.ProductUsecaseInterface
	CartUsecase    domain.CartUsecaseInterface
}

func GetAuthenticatedUser(ctx context.Context) *domain.User {
	user, authenticated := ctx.Value(auth.AuthenticatedUserContext).(*domain.User)
	if !authenticated {
		return nil
	}

	return user
}

func ConvertCart(user domain.User, cart domain.Cart) *gModel.Cart {
	var (
		cartItems []*gModel.Product
		itemsQty  map[uint64]int = map[uint64]int{}
	)

	for _, cartItem := range cart.Items {
		itemsQty[cartItem.ProductID] = cartItem.Qty
	}

	for _, cartItem := range cart.Products {
		cartItems = append(cartItems, &gModel.Product{
			ID:   int(cartItem.ID),
			Sku:  cartItem.SKU,
			Name: cartItem.Name,
			Price: &gModel.ProductPrice{
				OriginalPrice: cartItem.Price * float64(itemsQty[cartItem.ID]),
				Currency:      "$",
			},
			Qty: itemsQty[cartItem.ID],
		})
	}

	return &gModel.Cart{
		ID: int(cart.ID),
		Owner: &gModel.User{
			ID:       int(user.ID),
			Username: user.Username,
		},
		Products: cartItems,
	}
}
