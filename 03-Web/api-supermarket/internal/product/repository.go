package product

import (
	"api-supermarket/internal/domain"
)

type ProductRepository struct {
	productsDB []domain.Product
}

func NewProductRepository(prods []domain.Product) ProductRepository {
	return ProductRepository{prods}
}

func (r ProductRepository) GetAllProducts() []domain.Product {
	return r.productsDB
}

func (r ProductRepository) GetById(id int64) domain.Product {
	return r.productsDB[id-1]
}
