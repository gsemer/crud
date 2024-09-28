package persistence

import (
	"context"
	"crud/domain"
	"log"

	"github.com/arangodb/go-driver"
)

type ProductRepository struct {
	db         driver.Database
	collection driver.Collection
}

func NewProductRepository(db driver.Database, collection driver.Collection) *ProductRepository {
	return &ProductRepository{db: db, collection: collection}
}

func (pr ProductRepository) GetProducts() ([]domain.Product, error) {
	query := "FOR product IN products RETURN product"
	cursor, err := pr.db.Query(context.Background(), query, nil)
	if err != nil {
		log.Fatalf("Failed to execute query")
		return nil, err
	}
	defer cursor.Close()

	var products []domain.Product
	for {
		var product domain.Product
		meta, err := cursor.ReadDocument(context.Background(), &product)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Failed to read document")
			return nil, err
		}

		product = domain.Product{
			ID:       string(meta.ID),
			Key:      meta.Key,
			Name:     product.Name,
			Price:    product.Price,
			Category: product.Category,
			Quality:  product.Quality,
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr ProductRepository) GetProduct(key string) (domain.Product, error) {
	var product domain.Product

	meta, err := pr.collection.ReadDocument(context.Background(), key, &product)
	if err != nil {
		log.Fatalf("Failed to read document")
		return domain.Product{}, err
	}

	product = domain.Product{
		ID:       string(meta.ID),
		Key:      meta.Key,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
		Quality:  product.Quality,
	}
	return product, nil
}
