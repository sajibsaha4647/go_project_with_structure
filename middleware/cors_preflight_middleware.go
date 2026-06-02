package middleware

import (
	"net/http"
)
func CorsAndPreflightMiddleware(mux http.Handler) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type , sajibsaha")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.WriteHeader(http.StatusAccepted)

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)

	}

	handler := http.HandlerFunc(handleAllReq)

	return handler

}