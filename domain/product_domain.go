package domain

type ProductList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       string `json:"price" db:"price"`
	ImageUrl    string `json:"imageUrl" db:"image_url"`
}
