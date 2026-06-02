package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommerce/model"
	"ecommerce/utils"
)

func HellowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This was first step")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am software engineer")
}

var productList []model.ProductList

func GetProduct(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	customHeader := r.Header.Get("sajibsaha")
	fmt.Println(customHeader)

	response := model.Response{
		Message: "Data fetch successfully",
		Status:  http.StatusOK,
		Data:    productList,
	}

	utils.SendResponse(w, response)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "POST" {
		http.Error(w, "Please go with post method", 400)
		return
	}

	utils.HandlePreflightReq(w, r)

	var newProduct model.ProductList
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	newProduct.Id = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(http.StatusCreated)

	response := model.Response{
		Message: "Successfully created",
		Status:  http.StatusCreated,
		Data:    newProduct,
	}
	utils.SendResponse(w, response)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {

	productid := r.PathValue("id")

	id, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	for _, product := range productList {

		if id == product.Id {

			response := model.Response{
				Message: "Successfully fetched",
				Status:  http.StatusOK,
				Data:    product,
			}

			utils.SendResponse(w, response)
			return
		}
	}

	// only runs if no product found
	response := model.Response{
		Message: "did not find product",
		Status:  http.StatusNotFound,
		Data:    nil,
	}

	utils.SendResponse(w, response)

	fmt.Println("checking here id", id)
}

func init() {
	productList = append(productList, []model.ProductList{
		{Id: 1, Title: "Orange", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 2, Title: "Orange2", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 3, Title: "Orange3", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
		{Id: 4, Title: "Orange4", Description: "This was good", Price: "56", ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Oranges_-_whole-halved-segment.jpg/500px-Oranges_-_whole-halved-segment.jpg"},
	}...)
}
