package product

import (
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) productRepository {
	return productRepository{
		db: db,
	}
}

func (p *productRepository) Store(product domain.ProductList) (*domain.ProductList, error) {
	query := `INSERT INTO products (title, description, price, image_url) VALUES (:title, :description, :price, :image_url) RETURNING id`
	var storedProduct domain.ProductList
	err := p.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl).Scan(&storedProduct.Id)
	if err != nil {
		return nil, err
	}

	return &storedProduct, nil
}
func (p *productRepository) GetAllProducts() ([]domain.ProductList, error) {
	var products []domain.ProductList
	query := `SELECT id, title, description, price, image_url FROM products`
	err := p.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (p *productRepository) GetProductById(id int) (*domain.ProductList, error) {
	var product domain.ProductList
	query := `SELECT id, title, description, price, image_url FROM products WHERE id = ?`
	err := p.db.QueryRow(query, id).Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.ImageUrl)
	if err != nil {
		return &domain.ProductList{}, err
	}
	return &product, nil
}

// 		return product
// 	}
// }
// return model.ProductList{}

func (p *productRepository) UpdateProductById(id int, updatedProduct domain.ProductList) (*domain.ProductList, error) {
	// for i, product := range p.productilst {
	// 	if id == product.Id {
	// 		p.productilst[i] = updatedProduct
	// 		return updatedProduct
	// 	}
	// }
	// return model.ProductList{}
	query := `UPDATE products SET title = ?, description = ?, price = ?, image_url = ? WHERE id = ?`
	_, err := p.db.Exec(query, updatedProduct.Title, updatedProduct.Description, updatedProduct.Price, updatedProduct.ImageUrl, id)
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}
func (p *productRepository) DeleteProductById(id int) (bool, error) {
	// for i, product := range p.productilst {
	// 	if id == product.Id {
	// 		p.productilst = append(p.productilst[:i], p.productilst[i+1:]...)
	// 		return true
	// 	}
	// }
	// return false
	query := `DELETE FROM products WHERE id = ?`
	result, err := p.db.Exec(query, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
