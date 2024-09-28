package services

import "crud/domain"

type ProductService struct {
	pr domain.ProductRepository
}

func NewProductService(pr domain.ProductRepository) *ProductService {
	return &ProductService{pr: pr}
}

func (ps ProductService) GetProducts() ([]domain.Product, error) {
	products, err := ps.pr.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
