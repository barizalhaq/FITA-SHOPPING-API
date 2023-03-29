package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"net/http"

	"github.com/barizalhaq/fita_shopping_api/graph"
	"github.com/barizalhaq/fita_shopping_api/graph/gModel"
	"github.com/barizalhaq/fita_shopping_api/user/delivery/auth"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input gModel.RegisterInput) (*gModel.RegisterResponse, error) {
	user, err := r.UserUsecase.Register(input)
	if err != nil {
		return nil, err
	}

	return &gModel.RegisterResponse{
		User: &gModel.User{
			ID:       int(user.ID),
			Username: user.Username,
		},
	}, nil
}

// EnCart is the resolver for the enCart field.
func (r *mutationResolver) EnCart(ctx context.Context, input gModel.EncartInput) (*gModel.Cart, error) {
	authenticatedUser := GetAuthenticatedUser(ctx)
	if authenticatedUser == nil {
		return nil, gqlerror.Errorf("Unauthorized! Please Register or Sign In first!")
	}

	cart, err := r.CartUsecase.AddProducts(*authenticatedUser, input)
	if err != nil {
		return nil, err
	}

	return ConvertCart(*authenticatedUser, *cart), nil
}

// DeCart is the resolver for the deCart field.
func (r *mutationResolver) DeCart(ctx context.Context, input gModel.DecartInput) (*gModel.Cart, error) {
	authenticatedUser := GetAuthenticatedUser(ctx)
	if authenticatedUser == nil {
		return nil, gqlerror.Errorf("Unauthorized! Please Register or Sign In first!")
	}

	cart, err := r.CartUsecase.Decart(*authenticatedUser, input)
	if err != nil {
		return nil, err
	}

	return ConvertCart(*authenticatedUser, *cart), nil
}

// DecreaseCartProductQty is the resolver for the decreaseCartProductQty field.
func (r *mutationResolver) DecreaseCartProductQty(ctx context.Context, input gModel.DecreaseCartProductQtyInput) (*gModel.Cart, error) {
	authenticatedUser := GetAuthenticatedUser(ctx)
	if authenticatedUser == nil {
		return nil, gqlerror.Errorf("Unauthorized! Please Register or Sign In first!")
	}

	cart, err := r.CartUsecase.DecreaseQty(*authenticatedUser, input)
	if err != nil {
		return nil, err
	}

	return ConvertCart(*authenticatedUser, *cart), nil
}

// Checkout is the resolver for the checkout field.
func (r *mutationResolver) Checkout(ctx context.Context) (*gModel.Invoice, error) {
	authenticatedUser := GetAuthenticatedUser(ctx)
	if authenticatedUser == nil {
		return nil, gqlerror.Errorf("Unauthorized! Please Register or Sign In first!")
	}

	invoice, err := r.CartUsecase.Checkout(*authenticatedUser)
	if err != nil {
		return nil, err
	}

	cart := ConvertCart(*authenticatedUser, invoice.Cart)

	return &gModel.Invoice{
		Cart: cart,
		TotalPrice: &gModel.InvoicePrice{
			OriginalTotalPrice: invoice.TotalPrice.OriginalTotalPrice,
			ActualCurrentPrice: invoice.TotalPrice.ActualTotalPrice,
			PriceDiscount:      invoice.TotalPrice.Discount,
			Currency:           invoice.TotalPrice.Currency,
		},
	}, nil
}

// Authenticate is the resolver for the authenticate field.
func (r *queryResolver) Authenticate(ctx context.Context, input gModel.AuthenticateInput) (*gModel.AuthenticateResponse, error) {
	signaturedToken, err := r.UserUsecase.Authenticate(input)
	if err != nil {
		return nil, err
	}

	authCookieAcc := auth.GetCookieAccess(ctx)
	authCookieAcc.GinContext.SetSameSite(http.SameSiteLaxMode)
	authCookieAcc.GinContext.SetCookie(auth.AuthorizationCookieKey, signaturedToken, 3600, "", "", false, true)

	return &gModel.AuthenticateResponse{
		Authenticated: true,
	}, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*gModel.Product, error) {
	products, err := r.ProductUsecase.ListProducts()
	if err != nil {
		return nil, err
	}

	var gModelProducts []*gModel.Product

	for _, product := range products {
		gModelProduct := &gModel.Product{
			ID:   int(product.ID),
			Sku:  product.SKU,
			Name: product.Name,
			Price: &gModel.ProductPrice{
				OriginalPrice: product.Price,
				Currency:      "$",
			},
			Qty: product.Qty,
		}

		gModelProducts = append(gModelProducts, gModelProduct)
	}

	return gModelProducts, nil
}

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context) (*gModel.Cart, error) {
	authenticatedUser := GetAuthenticatedUser(ctx)
	if authenticatedUser == nil {
		return nil, gqlerror.Errorf("Unauthorized! Please Register or Sign In first!")
	}

	cart, err := r.CartUsecase.ViewCart(*authenticatedUser)
	if err != nil {
		return nil, err
	}

	return ConvertCart(*authenticatedUser, *cart), nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) RemoveCart(ctx context.Context) (*gModel.Cart, error) {
	panic(fmt.Errorf("not implemented: RemoveCart - removeCart"))
}