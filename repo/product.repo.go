package repo

import (
	"ecommerce/model"
	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	Store(model.ProductList) (*model.ProductList, error)
	GetAllProducts() ([]model.ProductList, error)
	GetProductById(id int) (model.ProductList, error)
	UpdateProductById(id int, updatedProduct model.ProductList) (*model.ProductList, error)
	DeleteProductById(id int) (bool, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Store(product model.ProductList) (*model.ProductList, error) {
	query := `INSERT INTO products (title, description, price, image_url) VALUES (:title, :description, :price, :image_url) RETURNING id`
	var storedProduct model.ProductList
	err := p.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl).Scan(&storedProduct.Id)
	if err != nil {
		return nil, err
	}

	return &storedProduct, nil
}
func (p *productRepo) GetAllProducts() ([]model.ProductList, error) {
	var products []model.ProductList
	query := `SELECT id, title, description, price, image_url FROM products`
	err := p.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (p *productRepo) GetProductById(id int) (model.ProductList, error) {
	var product model.ProductList
	query := `SELECT id, title, description, price, image_url FROM products WHERE id = ?`
	err := p.db.QueryRow(query, id).Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.ImageUrl)
	if err != nil {
		return model.ProductList{}, err
	}
	return product, nil
}
	// 		return product
	// 	}
	// }
	// return model.ProductList{}
	

func (p *productRepo) UpdateProductById(id int, updatedProduct model.ProductList) (*model.ProductList, error) {
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
func (p *productRepo) DeleteProductById(id int) (bool, error) {
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

func generateProducts(p *productRepo) {
	// p.productilst = append(p.productilst, []model.ProductList{
	// 	{Id: 1, Title: "Orange", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	// 	{Id: 2, Title: "Orange2", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	// 	{Id: 3, Title: "Orange3", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	// 	{Id: 4, Title: "Orange4", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	// }...)
}
