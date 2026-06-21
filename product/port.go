package product

import "ecommerce/domain"

type ProductRepo interface {
	Store(domain.ProductList) (*domain.ProductList, error)
	GetAllProducts() ([]domain.ProductList, error)
	GetProductById(id int) (*domain.ProductList, error)
	UpdateProductById(id int, updatedProduct domain.ProductList) (*domain.ProductList, error)
	DeleteProductById(id int) (bool, error)
}