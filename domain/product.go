package domain

type Product struct {
	ID    uint64
	SKU   string
	Name  string
	Price float64
	Qty   int
}

type ProductUsecaseInterface interface {
	ListProducts() ([]Product, error)
}

type ProductRepositoryInterface interface {
	GetProducts() ([]Product, error)
	GetProductByID(id uint64) (*Product, error)
}
