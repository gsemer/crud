package domain

type Product struct {
	ID       string `json:"_id,omitempty"`
	Key      string `json:"_key,omitempty"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Category string `json:"category"`
	Quality  string `json:"quality"`
}
