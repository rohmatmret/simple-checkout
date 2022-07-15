package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/simple-store/domain"
)

type storeHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewStoreHandler(r chi.Router, productUsecase domain.ProductUsecase) *storeHandler {
	return &storeHandler{
		ProductUsecase: productUsecase,
	}
}

// Set new Data product
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
func (h *storeHandler) Save(w http.ResponseWriter, r *http.Request) {
	var cart domain.RequestCart
	products := setNewProduct()
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	total := h.ProductUsecase.Save(cart.Name, products)
	s := fmt.Sprintf("%s%.2f", "$", total)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
	return
}
