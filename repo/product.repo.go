package repo

import "ecommerce/model"

type ProductRepo interface {
	Store(model.ProductList) []model.ProductList
	GetAllProducts() []model.ProductList
	GetProductById(id int) model.ProductList
	UpdateProductById(id int, updatedProduct model.ProductList) model.ProductList
	DeleteProductById(id int) bool
}

type productRepo struct {
	productilst []model.ProductList
}

func NewProductRepo() ProductRepo {
	return &productRepo{}
}


func (p *productRepo) Store(product model.ProductList) []model.ProductList {
	p.productilst = append(p.productilst, product)
	return p.productilst
}
func (p *productRepo) GetAllProducts() []model.ProductList {
	return p.productilst
}
func (p *productRepo) GetProductById(id int) model.ProductList {
	for _, product := range p.productilst {
		if id == product.Id {
			return product
		}
	}
	return model.ProductList{}
}
func (p *productRepo) UpdateProductById(id int, updatedProduct model.ProductList) model.ProductList {
	for i, product := range p.productilst {
		if id == product.Id {
			p.productilst[i] = updatedProduct
			return updatedProduct
		}
	}
	return model.ProductList{}
}
func (p *productRepo) DeleteProductById(id int) bool {
	for i, product := range p.productilst {
		if id == product.Id {
			p.productilst = append(p.productilst[:i], p.productilst[i+1:]...)
			return true
		}
	}
	return false
}

func generateProducts( p *productRepo) {
	p.productilst = append(p.productilst, []model.ProductList{
		{Id: 1, Title: "Orange", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 2, Title: "Orange2", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 3, Title: "Orange3", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 4, Title: "Orange4", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	}...)
}
