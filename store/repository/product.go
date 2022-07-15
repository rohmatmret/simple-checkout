package repository

import (
	"errors"
	"fmt"
	"math"

	"github.com/simple-store/domain"
)

type Stores struct {
}

func (s Stores) GetProduct(productName string, dataProducts []domain.Product) (domain.Product, error) {
	for _, p := range dataProducts {
		if productName == p.Name {
			return p, nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}

func getFreeItems(product domain.Stores) domain.Product {
	if product.Name == "Macbook Pro" {
		return domain.Product{
			SKU:      "234234",
			Name:     "Raspberry Pi B",
			Quantity: -1,
			Price:    -30.00,
		}
	}
	return domain.Product{}
}

func findIntemp(product domain.Stores, temp []domain.Stores) bool {
	for _, p := range temp {
		if p.Name == product.Name {
			return true
		}
	}
	return false
}

func DescribeOrders(product []domain.Stores) []domain.Stores {
	temp := []domain.Stores{}
	for _, p := range product {
		if findIntemp(p, temp) == false {
			temp = append(temp, domain.Stores{
				Name:  p.Name,
				Qty:   p.Qty,
				Price: p.Price,
			})
			if free := getFreeItems(p); free.Name != "" {
				temp = append(temp, domain.Stores{
					Name:  free.Name,
					Qty:   -1,
					Price: free.Price,
				})
			}
		} else {
			for n, p2 := range temp {
				if p2.Name == p.Name {
					temp[n].Qty += p.Qty
				}
			}

		}
	}
	return temp
}

func CalcPrice(product []domain.Stores) float64 {
	var total float64
	for _, p := range product {
		if p.Qty == 3 && p.Name == "Google Home" {
			total += p.Price * float64(p.Qty)
			total -= p.Price
		} else if p.Qty == 3 && p.Name == "Alexa Speaker" {
			tempPrice := p.Price * float64(p.Qty)
			discount := tempPrice * 0.1
			total += tempPrice - discount
		} else {
			fmt.Println("Total price: ", p.Name, p.Price*float64(p.Qty))
			total += p.Price * float64(p.Qty)
		}
	}
	return total
}

func (s Stores) Save(cart []domain.Cart, products []domain.Product) float64 {
	var validProducts []domain.Stores
	for _, p := range cart {
		product, err := s.GetProduct(p.Name, products)
		if err == nil {
			validProducts = append(validProducts, domain.Stores{
				Name:  product.Name,
				Qty:   1,
				Price: product.Price,
			})
		}
	}

	if len(validProducts) > 0 {
		cart := DescribeOrders(validProducts)
		finalPrice := CalcPrice(cart)
		return math.Round(finalPrice*100) / 100
	}
	return 0.0
}

func NewStore() domain.ProductsRepository {
	return Stores{}
}
