package controller

import (
	"ecommerce/model"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "POST" {
		http.Error(w, "Please go with post method", 400)
		return
	}
	var userLogin model.UserLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userLogin)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	isValid := model.ValidateLogin(userLogin)
	fmt.Println("Login attempt for email:", userLogin.Email, "Valid:", isValid)
	if !isValid {
		response := model.Response{
			Message: "Login Invalid credentials",
			Status:  400,
			Data: model.UserLogin{
				Email: userLogin.Email,
				Token: "",
			},
		}
		utils.SendResponse(w, response)
		return
	}

	user := model.FindUserByEmail(userLogin.Email)

	token, _ := utils.CreateJWT(utils.Payload{
		Sub:      strconv.Itoa(user.Id),
		Name:     user.Name,
		Email:    user.Email,
		UserType: "Owner",
		Iat:      time.Now().Unix(),
		Exp:      time.Now().Add(24 * time.Hour).Unix(),
	})

	response := model.Response{
		Message: "Login successful",
		Status:  http.StatusOK,
		Data:    token,
	}
	utils.SendResponse(w, response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	users := model.GetAllUsers()

	response := model.Response{
		Message: "Data fetch successfully",
		Status:  http.StatusOK,
		Data:    users,
	}
	utils.SendResponse(w, response)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "POST" {
		http.Error(w, "Please go with post method", 400)
		return
	}

	var user model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	model.StoreUser(user)

	response := model.Response{
		Message: "User created successfully",
		Status:  http.StatusCreated,
		Data:    user,
	}
	utils.SendResponse(w, response)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	productid := r.PathValue("id")
	id, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user := model.GetUserById(id)
	if user.Id == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := model.Response{
		Message: "User found",
		Status:  http.StatusOK,
		Data:    user,
	}
	utils.SendResponse(w, response)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "PUT" {
		http.Error(w, "Please go with put method", 400)
		return
	}
	utils.HandlePreflightReq(w, r)
	userId := r.PathValue("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	var updatedUser model.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	updatedUser.Id = id
	model.UpdateUser(id, updatedUser)

	response := model.Response{
		Message: "User updated successfully",
		Status:  http.StatusOK,
		Data:    updatedUser,
	}
	utils.SendResponse(w, response)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "DELETE" {
		http.Error(w, "Please go with delete method", 400)
		return
	}
	utils.HandlePreflightReq(w, r)
	userId := r.PathValue("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	deleted := model.DeleteUser(id)
	if !deleted {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	response := model.Response{
		Message: "User deleted successfully",
		Status:  http.StatusOK,
		Data:    nil,
	}
	utils.SendResponse(w, response)

}
