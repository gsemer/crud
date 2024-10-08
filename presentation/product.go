package presentation

import (
	"crud/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ps domain.ProductService
}

func NewProductHandler(ps domain.ProductService) *ProductHandler {
	return &ProductHandler{ps: ps}
}

func (ph ProductHandler) GetProducts(writer http.ResponseWriter, request *http.Request) {
	p, ok1 := request.URL.Query()["page"]
	l, ok2 := request.URL.Query()["limit"]
	var page, limit string
	if ok1 && ok2 {
		page = p[0]
		limit = l[0]
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Please make sure that you use all required query parameters: `page`, `limit`"))
		return
	}

	pageToInt, err := strconv.Atoi(page)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Please make sure that `page` is an integer"))
		return
	}
	limitToInt, err := strconv.Atoi(limit)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Please make sure that `limit` is an integer"))
		return
	}

	products, err := ph.ps.GetProducts(pageToInt, limitToInt)
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
}

func (ph ProductHandler) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		result, _ := json.Marshal(err)
		writer.Write(result)
		return
	}
	defer request.Body.Close()

	var product domain.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		result, _ := json.Marshal(err)
		writer.Write(result)
		return
	}

	product, err = ph.ps.CreateProduct(product)
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
}
