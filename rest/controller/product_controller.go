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
		Data:    model.GetAllProducts(),
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

	newProduct.Id = len(model.GetAllProducts()) + 1
	model.Store(newProduct)

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

	for _, product := range model.GetAllProducts() {

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

func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "PUT" {
		http.Error(w, "Please go with put method", 400)
		return
	}

	utils.HandlePreflightReq(w, r)

	productid := r.PathValue("id")

	id, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	var updatedProduct model.ProductList
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	updatedProduct.Id = id

	result := model.UpdateProductById(id, updatedProduct)

	if result.Id == 0 {
		response := model.Response{
			Message: "did not find product",
			Status:  http.StatusNotFound,
			Data:    nil,
		}
		utils.SendResponse(w, response)
		return
	}

	response := model.Response{
		Message: "Successfully updated",
		Status:  http.StatusOK,
		Data:    result,
	}
	utils.SendResponse(w, response)

}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "DELETE" {
		http.Error(w, "Please go with delete method", 400)
		return
	}

	utils.HandlePreflightReq(w, r)

	productid := r.PathValue("id")

	id, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	isDeleted := model.DeleteProductById(id)

	if !isDeleted {
		response := model.Response{
			Message: "did not find product",
			Status:  http.StatusNotFound,
			Data:    nil,
		}
		utils.SendResponse(w, response)
		return
	}

	response := model.Response{
		Message: "Successfully deleted",
		Status:  http.StatusOK,
		Data:    nil,
	}
	utils.SendResponse(w, response)
}
