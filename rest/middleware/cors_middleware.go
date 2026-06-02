package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type , sajibsaha")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		next.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleCors)

	return handler

}