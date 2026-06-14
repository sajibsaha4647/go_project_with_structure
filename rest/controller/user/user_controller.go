package user

import (
	"ecommerce/model"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "POST" {
		http.Error(w, "Please go with post method", 400)
		return
	}
	var userLogin model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userLogin)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	isValid, err := h.loginRepo.ValidateLogin(userLogin)
	if err != nil {
		http.Error(w, "Error validating login", 400)
		return
	}
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

	user, err := h.loginRepo.FindUserByEmail(userLogin.Email)
	if err != nil {
		http.Error(w, "Error finding user", 400)
		return
	}

	token, _ := utils.CreateJWT(os.Getenv("JWT_SECRET"), utils.Payload{
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

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	response := model.Response{
		Message: "Data fetch successfully",
		Status:  http.StatusOK,
		Data:    users,
	}
	utils.SendResponse(w, response)

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	_, err = h.userRepo.Store(user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	response := model.Response{
		Message: "User created successfully",
		Status:  http.StatusCreated,
		Data:    user,
	}
	utils.SendResponse(w, response)

}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	userid := r.PathValue("id")
	id, err := strconv.Atoi(userid)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user := h.userRepo.GetUserById(id)
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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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
	h.userRepo.UpdateUserById(id, updatedUser)

	response := model.Response{
		Message: "User updated successfully",
		Status:  http.StatusOK,
		Data:    updatedUser,
	}
	utils.SendResponse(w, response)

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	deleted := h.userRepo.DeleteUserById(id)
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
