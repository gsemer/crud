package presentation

import (
	"crud/app/fakes"
	"crud/domain"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetProducts(t *testing.T) {
	ps := &fakes.FakeProductService{}
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
	ps.GetProductsReturns([]domain.Product{products[0]}, nil)

	myHandlerFunc := ProductHandler{ps: ps}

	values := url.Values{}
	values.Set("page", "1")
	values.Set("limit", "1")

	request := httptest.NewRequest(http.MethodGet, "/products?"+values.Encode(), nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/products", myHandlerFunc.GetProducts)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code is not %v, but is %v", http.StatusOK, response.Code)
		return
	}

	productsList, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	var result []domain.Product
	err = json.Unmarshal([]byte(productsList), &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result, []domain.Product{products[0]}) {
		t.Error("Not the expected output")
		return
	}
}

func TestGetProducts_FAIL(t *testing.T) {
	ps := &fakes.FakeProductService{}
	ps.GetProductsReturns(nil, errors.New("An error occurred during the execution"))

	myHandlerFunc := ProductHandler{ps: ps}

	server := httptest.NewServer(http.HandlerFunc(myHandlerFunc.GetProducts))
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
		return
	}
	res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %v but got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestGetProducts_FAIL2(t *testing.T) {
	ps := &fakes.FakeProductService{}
	ps.GetProductsReturns(nil, nil)

	myHandlerFunc := ProductHandler{ps: ps}

	values := url.Values{}
	values.Set("page", "1.5")
	values.Set("limit", "1")

	request := httptest.NewRequest(http.MethodGet, "/products?"+values.Encode(), nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/products", myHandlerFunc.GetProducts)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Status code is not %v, but is %v", http.StatusBadRequest, response.Code)
		return
	}
}
