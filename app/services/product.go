package services

import "crud/domain"

type ProductService struct {
	pr domain.ProductRepository
}

func NewProductService(pr domain.ProductRepository) *ProductService {
	return &ProductService{pr: pr}
}

func (ps ProductService) GetProducts(page int, limit int) ([]domain.Product, error) {
	products, err := ps.pr.GetProducts(page, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps ProductService) GetProduct(key string) (domain.Product, error) {
	product, err := ps.pr.GetProduct(key)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (ps ProductService) CreateProduct(product domain.Product) (domain.Product, error) {
	product, err := ps.pr.CreateProduct(product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
