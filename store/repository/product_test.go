package repository

import (
	"testing"

	"github.com/simple-store/domain"
)

func setNewProduct() []domain.Product {
	return []domain.Product{
		{
			SKU:      "120P90",
			Name:     "Google Home",
			Price:    49.99,
			Quantity: 10,
		},
		{
			SKU:      "43N23P",
			Name:     "Macbook Pro",
			Price:    5399.99,
			Quantity: 5,
		},
		{
			SKU:      "A304SD",
			Name:     "Alexa Speaker",
			Price:    109.50,
			Quantity: 10,
		},
		{
			SKU:      "234234",
			Name:     "Raspberry Pi B",
			Price:    30.00,
			Quantity: 2,
		},
	}
}
func TestStores_Save(t *testing.T) {
	type args struct {
		cart    []domain.Cart
		product []domain.Product
	}
	tests := []struct {
		name string
		s    Stores
		args args
		want float64
	}{
		{
			name: "TestSave Macbook Pro with free Raspberry Pi B",
			s:    Stores{},
			args: args{
				cart: []domain.Cart{
					{
						Name: "Macbook Pro",
					},
					{
						Name: "Raspberry Pi B",
					},
				},
				product: setNewProduct(),
			},
			want: 5399.99,
		},
		{
			name: "TestSave",
			s:    Stores{},
			args: args{
				cart: []domain.Cart{
					{
						Name: "Google Home",
					},
					{
						Name: "Google Home",
					},
					{
						Name: "Google Home",
					},
				},
				product: setNewProduct(),
			},
			want: 99.98,
		},
		{
			name: "TestSave",
			s:    Stores{},
			args: args{
				cart: []domain.Cart{
					{
						Name: "Alexa Speaker",
					},
					{
						Name: "Alexa Speaker",
					},
					{
						Name: "Alexa Speaker",
					},
				},
				product: setNewProduct(),
			},
			want: 295.65,
		},
		{
			name: "TestSave",
			s:    Stores{},
			args: args{
				cart: []domain.Cart{
					{
						Name: "Alexa Speaker",
					},
					{
						Name: "Alexa Speaker",
					},
					{
						Name: "Alexa Speaker",
					},
					{
						Name: "Macbook Pro",
					},
					{
						Name: "Raspberry Pi B",
					},
				},
				product: setNewProduct(),
			},
			want: 5695.64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stores{}
			if got := s.Save(tt.args.cart, tt.args.product); got != tt.want {
				t.Errorf("Stores.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
