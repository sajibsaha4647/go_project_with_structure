package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommerce/model"
	"ecommerce/utils"
)

func (h *Handler) HellowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This was first step")
}

func (h *Handler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am software engineer")
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	customHeader := r.Header.Get("sajibsaha")
	fmt.Println(customHeader)

	products, err := h.repo.GetAllProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	response := model.Response{
		Message: "Data fetch successfully",
		Status:  http.StatusOK,
		Data:    products,
	}

	utils.SendResponse(w, response)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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
	products, error := h.repo.GetAllProducts()
	if error != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	newProduct.Id = len(products) + 1
	h.repo.Store(newProduct)

	w.WriteHeader(http.StatusCreated)

	response := model.Response{
		Message: "Successfully created",
		Status:  http.StatusCreated,
		Data:    newProduct,
	}
	utils.SendResponse(w, response)
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {

	productid := r.PathValue("id")

	id, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	products, err := h.repo.GetAllProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	for _, product := range products {

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

func (h *Handler) UpdateProductById(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.repo.UpdateProductById(id, updatedProduct)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

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

func (h *Handler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
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

	isDeleted, err := h.repo.DeleteProductById(id)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

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
