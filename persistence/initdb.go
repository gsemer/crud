package persistence

import (
	"context"
	"crud/domain"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/arangodb/go-driver"
)

type ClientRepository struct {
	client driver.Client
}

func NewClientRepository(client driver.Client) *ClientRepository {
	return &ClientRepository{client: client}
}

type DatabaseRepository struct {
	db driver.Database
}

func NewDatabaseRepository(db driver.Database) *DatabaseRepository {
	return &DatabaseRepository{db: db}
}

type CollectionRepository struct {
	collection driver.Collection
}

func NewCollectionRepository(collection driver.Collection) *CollectionRepository {
	return &CollectionRepository{collection: collection}
}

func (cr ClientRepository) GetOrCreateDB() (driver.Database, error) {
	db_exists, err := cr.client.DatabaseExists(context.Background(), "crud")
	if err != nil {
		log.Printf("Failed to check if database exists")
	}
	var db driver.Database
	if !db_exists {
		db, err = cr.client.CreateDatabase(context.Background(), "crud", nil)
		if err != nil {
			log.Printf("Failed to create database")
			return nil, err
		}
	} else {
		db, err = cr.client.Database(context.Background(), "crud")
		if err != nil {
			log.Printf("Failed to open database")
			return nil, err
		}
	}
	return db, nil
}

func (dr DatabaseRepository) GetOrCreateCollection() (driver.Collection, error) {
	collection, err := dr.db.Collection(context.Background(), "products")
	if driver.IsNotFoundGeneral(err) {
		collection, err = dr.db.CreateCollection(context.Background(), "products", nil)
		if err != nil {
			log.Printf("Failed to create collection")
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return collection, nil
}

func (cr CollectionRepository) BulkCreate(n int64) ([]domain.Product, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Buld create 100 random products in the collection `products`
	qualities := []string{"A", "B", "C", "D", "E"}
	categories := []string{"Category I", "Category II", "Category III"}
	products := make([]domain.Product, n)
	for i := 0; i < 100; i++ {
		products[i] = domain.Product{
			Name:     fmt.Sprintf("Name%d", i+1),
			Price:    fmt.Sprintf("%d", rand.Intn(1000-99)+99),
			Category: categories[r.Intn(len(categories))],
			Quality:  qualities[r.Intn(len(qualities))],
		}
	}

	// Insert products into the collection
	var dbProducts []domain.Product
	for _, product := range products {
		meta, err := cr.collection.CreateDocument(context.Background(), product)
		if err != nil {
			log.Printf("An error occurred during bulk creation")
			return nil, err
		}
		product.Key = meta.Key
		product.ID = meta.ID.String()
		dbProducts = append(dbProducts, product)
	}

	return dbProducts, nil
}

func (cr CollectionRepository) IsEmpty() bool {
	numOfDocs, err := cr.collection.Count(context.Background())
	if err != nil {
		log.Printf("Failed to get document count")
	}
	return numOfDocs == 0
}
