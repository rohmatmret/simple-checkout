package usecase

import (
	"github.com/simple-store/domain"
	store "github.com/simple-store/store/repository"
)

type Product struct {
	ProductRepository store.Stores
}

func (p *Product) GetProduct(productName string, dataProduct []domain.Product) (domain.Product, error) {
	return p.ProductRepository.GetProduct(productName, dataProduct)
}

func (p *Product) Save(cart []domain.Cart, products []domain.Product) float64 {
	return p.ProductRepository.Save(cart, products)
}

func NewProductUsecase() Product {
	return Product{}
}
