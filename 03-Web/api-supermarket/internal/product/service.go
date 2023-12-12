package product

import (
	"api-supermarket/internal/domain"
)

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{repo}
}

func (s *ProductService) GetAllProducts() []domain.Product {
	// pasamanos hacia el controller porque aca no tenemos logica de negocios aun
	return s.repository.GetAllProducts()
}

func (s *ProductService) GetById(id int64) domain.Product {
	// pasamanos hacia el controller porque aca no tenemos logica de negocios aun
	return s.repository.GetById(id)
}

func (s *ProductService) GetByPrice(price float64) []domain.Product {
	// Aca ya hay logica de negocios, por eso hay mucho mas codigo.
	products := s.repository.GetAllProducts()
	queryResult := []domain.Product{}
	for i, prod := range products {
		if prod.Price >= price {
			queryResult = append(queryResult, products[i])
		}
	}
	return queryResult
}
