package domain

import "github.com/barizalhaq/fita_shopping_api/graph/gModel"

type Cart struct {
	ID       uint64
	UserID   uint64
	Products []Product  `gorm:"many2many:cart_items;"`
	Items    []CartItem `gorm:"-"`
}

type CartItem struct {
	CartID    uint64 `gorm:"primaryKey"`
	ProductID uint64 `gorm:"primaryKey"`
	Qty       int
}

type Invoice struct {
	InvoiceID  string
	Cart       Cart
	TotalPrice struct {
		OriginalTotalPrice float64
		ActualTotalPrice   float64
		Discount           float64
		Currency           string
	}
}

type CartUsecaseInterface interface {
	ViewCart(user User) (*Cart, error)
	AddProducts(user User, input gModel.EncartInput) (*Cart, error)
	Decart(user User, input gModel.DecartInput) (*Cart, error)
	DecreaseQty(user User, input gModel.DecreaseCartProductQtyInput) (*Cart, error)
	Checkout(user User) (*Invoice, error)
}

type CartRepositoryInterface interface {
	Create(user *User) (*Cart, error)
	View(user *User) (*Cart, error)
	AddProducts(user *User, items map[Product]int) (*Cart, error)
	RemoveProducts(user *User, productIds []int) (*Cart, error)
	SubtractProducts(user *User, items map[Product]int) (*Cart, error)
}
