package product

import (
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductNewRepo {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) Store(product domain.ProductList) (*domain.ProductList, error) {
	query := `INSERT INTO products (title, description, price, image_url) VALUES ($1, $2, $3, $4) RETURNING id`
	var storedProduct domain.ProductList
	err := p.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl).Scan(&storedProduct.Id)
	if err != nil {
		return nil, err
	}

	return &storedProduct, nil
}
func (p *productRepository) GetProduct(page, limit int) ([]domain.ProductList, error) {
	offset := (page - 1) * limit

	var products []domain.ProductList

	query := `
		SELECT id, title, description, price, image_url
		FROM products
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	err := p.db.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}
func (p *productRepository) GetProductById(id int) (*domain.ProductList, error) {
	var product domain.ProductList
	query := `SELECT id, title, description, price, image_url FROM products WHERE id = $1`
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
	query := `UPDATE products SET title = $1, description = $2, price = $3, image_url = $4 WHERE id = $5`
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
	query := `DELETE FROM products WHERE id = $1`
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

func (p *productRepository) RowCount() (int64, error) {
	var count int64
	query := `SELECT COUNT(*) FROM products`
	err := p.db.QueryRow(query).Scan(&count)	
if err != nil {
		return 0, err
	}
	return count, nil
}
