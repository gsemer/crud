## How to Run
- `docker network create crud-network`
- `docker-compose up --build`

## Examples
- `http://localhost:8000/products?page=2&limit=5` 
-> This endpoint returns 5 products which can be found on page 2.
- `http://localhost:8000/products/176`
-> This endpoint returns the product with id 176.
- `http://localhost:8000/products/create` 
-> This endpoint creates a new product entry in the database.