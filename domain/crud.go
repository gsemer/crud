package domain

type Product struct {
	ID       string `json:"_id,omitempty"`
	Key      string `json:"_key,omitempty"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Category string `json:"category"`
	Quality  string `json:"quality"`
}

type Error struct {
	Status      string
	Description string
}

type ProductService interface {
	GetProducts() ([]Product, error)
	GetProduct(key string) (Product, error)
	CreateProduct(Product) (Product, error)
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProduct(key string) (Product, error)
	CreateProduct(Product) (Product, error)
}
