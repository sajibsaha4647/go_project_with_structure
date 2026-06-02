package utils

import (
	"encoding/json"
	"net/http"
)


func SendResponse(w http.ResponseWriter, response interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}









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