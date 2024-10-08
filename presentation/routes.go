package presentation

import "crud/domain"

func CreateRoutes(service domain.ProductService) map[string]domain.RouteDefinition {
	ph := NewProductHandler(service)
	// Define a map
	return map[string]domain.RouteDefinition{
		"/products": {
			Methods:     []string{"GET"},
			HandlerFunc: ph.GetProducts,
		},
		"/products/{key}": {
			Methods:     []string{"GET"},
			HandlerFunc: ph.GetProduct,
		},
		"/products/create": {
			Methods:     []string{"POST"},
			HandlerFunc: ph.CreateProduct,
		},
	}
}
