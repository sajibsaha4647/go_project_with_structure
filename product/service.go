package product

import "ecommerce/domain"

type Service interface {
	GetProduct(page, limit int) ([]domain.ProductList, error)
	CreateProduct(product domain.ProductList) (*domain.ProductList, error)
	GetProductById(id int) (*domain.ProductList, error)
	UpdateProductById(id int, product domain.ProductList) (*domain.ProductList, error)
	DeleteProductById(id int) (bool, error)
	RowCount() (int64, error)
}



type service struct {
	repo ProductNewRepo
}

func NewService(repo ProductNewRepo) Service {
	return &service{repo: repo}
}

// CreateProduct implements [Service].
func (s *service) CreateProduct(product domain.ProductList) (*domain.ProductList, error) {
	return s.repo.Store(product)
}

// DeleteProductById implements [Service].
func (s *service) DeleteProductById(id int) (bool, error) {
	return s.repo.DeleteProductById(id)
}

// GetProduct implements [Service].
func (s *service) GetProduct(page, limit int) ([]domain.ProductList, error) {
	return s.repo.GetProduct(page, limit)
}

// GetProductById implements [Service].
func (s *service) GetProductById(id int) (*domain.ProductList, error) {
	return s.repo.GetProductById(id)
}

// UpdateProductById implements [Service].
func (s *service) UpdateProductById(id int, product domain.ProductList) (*domain.ProductList, error) {
	return s.repo.UpdateProductById(id, product)
}

func (s *service) RowCount() (int64, error) {
	return s.repo.RowCount()
}
