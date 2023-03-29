package domain

type Promo struct {
	SKU                string `json:"sku"`
	OnFree             bool   `json:"on_free"`
	OnDiscount         bool   `json:"on_discount"`
	PurchaseAmount     int    `json:"purchase_amount"`
	DiscountPercentage int    `json:"discount_percentage"`
	FreeProduct        struct {
		SKU string `json:"sku"`
	} `json:"free_product"`
}

type PromoUsecaseInterface interface {
	Promos() ([]Promo, error)
}

type PromoRepositoryInterface interface {
	GetPromos() ([]Promo, error)
}
