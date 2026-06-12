package utils

import (
	"encoding/json"
	"net/http"
)


func SendResponse(w http.ResponseWriter, response interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// func SendResponse(w http.ResponseWriter, response interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(response.Status)

// 	if err := json.NewEncoder(w).Encode(response); err != nil {
// 		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 	}
// }









func HandelCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type , sajibsaha")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
}

func HandlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}

func CorsAndPreflightMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(w, r)
	})
}