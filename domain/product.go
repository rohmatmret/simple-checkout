package domain

type Cart struct {
	Name string `json:"name" validate:"required"`
}
type RequestCart struct {
	Name []Cart `json:"Items" validate:"required"`
}
type Product struct {
	SKU      string
	Name     string
	Price    float64
	Quantity int
}
type Stores struct {
	Name  string
	Qty   int
	Price float64
}

type ProductsRepository interface {
	GetProduct(productName string, dataProduct []Product) (Product, error)
	Save(cart []Cart, product []Product) float64
}

type ProductUsecase interface {
	Save(cart []Cart, product []Product) float64
}

func NewProduct(sku string, name string, price float64, quantity int) Product {
	return Product{
		SKU:      sku,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}
