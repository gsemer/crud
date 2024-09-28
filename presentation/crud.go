package presentation

import (
	"crud/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ps domain.ProductService
}

func NewProductHandler(ps domain.ProductService) *ProductHandler {
	return &ProductHandler{ps: ps}
}

func (ph ProductHandler) GetProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := ph.ps.GetProducts()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err2 := domain.Error{
			Status:      "error",
			Description: err.Error(),
		}
		resultErr, _ := json.Marshal(err2)
		writer.Write(resultErr)
		return
	}
	writer.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(products)
	writer.Write(result)
	return
}

func (ph ProductHandler) GetProduct(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key, ok := vars["key"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		result, _ := json.Marshal("Key is missing")
		writer.Write(result)
		return
	}

	product, err := ph.ps.GetProduct(key)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err2 := domain.Error{
			Status:      "error",
			Description: err.Error(),
		}
		resultErr, _ := json.Marshal(err2)
		writer.Write(resultErr)
		return
	}
	writer.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(product)
	writer.Write(result)
	return
}
