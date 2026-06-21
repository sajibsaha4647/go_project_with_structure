package product

import "ecommerce/domain"

type ProductNewRepo interface {
	Store(domain.ProductList) (*domain.ProductList, error)
	GetProduct(page,limit int) ([]domain.ProductList, error)
	GetProductById(id int) (*domain.ProductList, error)
	RowCount() (int64, error)
	UpdateProductById(id int, updatedProduct domain.ProductList) (*domain.ProductList, error)
	DeleteProductById(id int) (bool, error)
}