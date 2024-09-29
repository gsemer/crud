package services

import (
	"crud/domain"
	"crud/persistence/fakes"
	"reflect"
	"testing"
)

func TestGetProducts(t *testing.T) {
	pr := &fakes.FakeProductRepository{}
	products := []domain.Product{
		domain.Product{
			ID:       "123e4567-e89b-12d3-a456-426614174000",
			Key:      "1234",
			Name:     "Product I",
			Price:    "12.5",
			Category: "Caterory II",
			Quality:  "D",
		},
		domain.Product{
			ID:       "123e5567-e89b-12d3-a456-426614174003",
			Key:      "1235",
			Name:     "Product II",
			Price:    "124.7",
			Category: "Caterory I",
			Quality:  "C",
		},
	}
	pr.GetProductsReturns([]domain.Product{products[0]}, nil)

	productService := ProductService{pr: pr}

	dbProducts, err := productService.GetProducts(1, 1)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(dbProducts, []domain.Product{products[0]}) {
		t.Errorf("It was expected to have %v as output,\n but we got %v instead", []domain.Product{products[0]}, dbProducts)
		return
	}
}
