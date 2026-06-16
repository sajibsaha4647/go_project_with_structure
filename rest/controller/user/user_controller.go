package user

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
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

	var req domain.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	user, valid, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Error validating login", 400)
		return
	}
	if !valid {
		utils.SendResponse(w, domain.Response{
			Message: "Invalid credentials",
			Status:  400,
			Data:    domain.UserLogin{Email: req.Email},
		})
		return
	}

	token, _ := utils.CreateJWT(os.Getenv("JWT_SECRET"), utils.Payload{
		Sub:      strconv.Itoa(user.Id),
		Name:     user.Name,
		Email:    user.Email,
		UserType: user.UserType,
		Iat:      time.Now().Unix(),
		Exp:      time.Now().Add(24 * time.Hour).Unix(),
	})

	utils.SendResponse(w, domain.Response{
		Message: "Login successful",
		Status:  http.StatusOK,
		Data:    token,
	})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, domain.Response{
		Message: "Data fetch successfully",
		Status:  http.StatusOK,
		Data:    users,
	})
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "POST" {
		http.Error(w, "Please go with post method", 400)
		return
	}

	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	created, err := h.service.CreateUser(u)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, domain.Response{
		Message: "User created successfully",
		Status:  http.StatusCreated,
		Data:    created,
	})
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "GET" {
		http.Error(w, "Please go with get method", 400)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	u := h.service.GetUserById(id)
	if u == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, domain.Response{
		Message: "User found",
		Status:  http.StatusOK,
		Data:    u,
	})
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "PUT" {
		http.Error(w, "Please go with put method", 400)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	result := h.service.UpdateUserById(id, u)
	utils.SendResponse(w, domain.Response{
		Message: "User updated successfully",
		Status:  http.StatusOK,
		Data:    result,
	})
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.HandelCors(w)
	if r.Method != "DELETE" {
		http.Error(w, "Please go with delete method", 400)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if !h.service.DeleteUserById(id) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, domain.Response{
		Message: "User deleted successfully",
		Status:  http.StatusOK,
		Data:    nil,
	})
}
