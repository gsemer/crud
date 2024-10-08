package main

import (
	"crud/app/services"
	"crud/persistence"
	"crud/presentation"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arangodb/go-driver"
	arangohttp "github.com/arangodb/go-driver/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server is up and running...")

	// connect to ArangoDB
	var client driver.Client
	for i := 0; i < 5; i++ {
		conn, err := arangohttp.NewConnection(
			arangohttp.ConnectionConfig{
				Endpoints: []string{"http://arangodb_db_container:8529"},
			},
		)
		if err != nil {
			log.Fatalf("Failed to create HTTP connection: %v", err)
		}
		client, err = driver.NewClient(
			driver.ClientConfig{
				Connection:     conn,
				Authentication: driver.BasicAuthentication("root", "rootpassword"),
			},
		)
		if err != nil {
			log.Fatalf("Failed to create the client: %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	// Create database
	clientRepo := persistence.NewClientRepository(client)
	db, err := clientRepo.GetOrCreateDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Create collection
	dbRepo := persistence.NewDatabaseRepository(db)
	collection, err := dbRepo.GetOrCreateCollection()
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Bulk create 100 documents in the collection
	collectionRepo := persistence.NewCollectionRepository(collection)
	if collectionRepo.IsEmpty() {
		_, err = collectionRepo.BulkCreate(100)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	// router
	r := mux.NewRouter()

	productRepository := persistence.NewProductRepository(db, collection)
	productService := services.NewProductService(productRepository)

	productRoutes := presentation.CreateRoutes(productService)
	for routePath, routeMethods := range productRoutes {
		fmt.Printf("adding %s route with methods %v\n", routePath, routeMethods.Methods)
		r.Handle(routePath, routeMethods.HandlerFunc).Methods(routeMethods.Methods...)
	}

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
