package model

var productList []ProductList

type ProductList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

func Store(p ProductList) []ProductList {
	productList = append(productList, p)
	return productList
}

func GetAllProducts() []ProductList {
	return productList
}

func GetProductById(id int) ProductList {
	for _, product := range productList {
		if id == product.Id {
			return product
		}
	}
	return ProductList{}
}

func UpdateProductById(id int, updatedProduct ProductList) ProductList {
	for i, product := range productList {
		if id == product.Id {
			productList[i] = updatedProduct
			return updatedProduct
		}
	}
	return ProductList{}
}

func DeleteProductById(id int) bool {
	for i, product := range productList {
		if id == product.Id {
			productList = append(productList[:i], productList[i+1:]...)
			return true
		}
	}
	return false
}

func init() {
	productList = append(productList, []ProductList{
		{Id: 1, Title: "Orange", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 2, Title: "Orange2", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 3, Title: "Orange3", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 4, Title: "Orange4", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	}...)
}
