package usecase

import (
	"errors"
	"log"
	"math"

	"github.com/barizalhaq/fita_shopping_api/domain"
)

func (u *cartUsecase) Checkout(user domain.User) (*domain.Invoice, error) {
	var (
		promoSKUMap                map[string]domain.Promo           = map[string]domain.Promo{}
		productsMap                map[uint64]domain.Product         = map[uint64]domain.Product{}
		productsMapSKU             map[string]map[string]interface{} = map[string]map[string]interface{}{}
		alreadyCounted             map[string]struct{}               = map[string]struct{}{}
		originalPrice, actualPrice float64
	)

	existingCart, err := u.cartRepo.View(&user)
	if err != nil {
		log.Fatalf("[cartUsecase.Checkout] cartRepo.View Error: %v", err)
		return nil, err
	}

	if existingCart == nil {
		return nil, errors.New("no cart available")
	}

	promos, err := u.promoRepo.GetPromos()
	if err != nil {
		log.Fatalf("[cartUsecase.Checkout] promoRepo.GetPromos Error: %v", err)
		return nil, err
	}

	for _, promo := range promos {
		promoSKUMap[promo.SKU] = promo
	}

	for _, cartProduct := range existingCart.Products {
		productsMap[cartProduct.ID] = cartProduct
		productsMapSKU[cartProduct.SKU] = map[string]interface{}{
			"product": cartProduct,
		}
	}

	for _, item := range existingCart.Items {
		productsMapSKU[productsMap[item.ProductID].SKU]["item"] = item
	}

	for _, cartItem := range existingCart.Items {
		product := productsMap[cartItem.ProductID]

		if _, ok := alreadyCounted[product.SKU]; ok {
			continue
		}

		originalPrice += (float64(cartItem.Qty) * product.Price)
		actualPrice += (float64(cartItem.Qty) * product.Price)
		if promo, ok := promoSKUMap[product.SKU]; ok {
			if promo.OnFree {
				if freeProduct, ok := productsMapSKU[promo.FreeProduct.SKU]["product"].(domain.Product); ok && cartItem.Qty >= promo.PurchaseAmount {
					if promo.SKU == promo.FreeProduct.SKU {
						freeProductAmount := cartItem.Qty / promo.PurchaseAmount
						actualPrice -= (freeProduct.Price * float64(freeProductAmount))

						continue
					}

					freeProductAmount := cartItem.Qty / promo.PurchaseAmount

					if freeCartItemQty := productsMapSKU[freeProduct.SKU]["item"].(domain.CartItem).Qty; freeCartItemQty >= freeProductAmount {
						actualPrice -= (freeProduct.Price * float64(freeProductAmount))
					} else {
						actualPrice -= (freeProduct.Price * float64(freeCartItemQty))
					}

					alreadyCounted[promo.SKU] = struct{}{}
				}
			}

			if promo.OnDiscount {
				if freeProduct, ok := productsMapSKU[promo.SKU]["product"].(domain.Product); ok && cartItem.Qty >= promo.PurchaseAmount {
					discountPrice := (float64(promo.DiscountPercentage) * float64(freeProduct.Price)) / 100
					actualPrice -= discountPrice * float64(cartItem.Qty)

					alreadyCounted[promo.SKU] = struct{}{}
				}
			}
		}
	}

	return &domain.Invoice{
		InvoiceID: "invoice_id",
		Cart:      *existingCart,
		TotalPrice: struct {
			OriginalTotalPrice float64
			ActualTotalPrice   float64
			Discount           float64
			Currency           string
		}{
			OriginalTotalPrice: roundFloat(originalPrice, 2),
			ActualTotalPrice:   roundFloat(actualPrice, 2),
			Discount:           roundFloat(originalPrice-actualPrice, 2),
			Currency:           "$",
		},
	}, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
