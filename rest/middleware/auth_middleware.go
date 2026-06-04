package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	header := r.Header.Get("Authorization")

	if header == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	headerArray := strings.Split(header, " ")

	if len(headerArray) != 2 || headerArray[0] != "Bearer" {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}

	accessToken := headerArray[1]

	fmt.Println("Access Token:", accessToken)
	tokenParts := strings.Split(accessToken, ".")

	if len(tokenParts) != 3 {
		http.Error(w, "Invalid JWT token format", http.StatusUnauthorized)
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	jwtSignature := tokenParts[2]

	fmt.Println("JWT Header:", jwtHeader)
	fmt.Println("JWT Payload:", jwtPayload)
	fmt.Println("JWT Signature:", jwtSignature)

	message := jwtHeader + "." + jwtPayload

	fmt.Println("Message to verify:", message)

	biteArraySecret := []byte(os.Getenv("JWT_SECRET"))
	biteArrayMessage := []byte(message)

	h := hmac.New(sha256.New, biteArraySecret)
	h.Write(biteArrayMessage)

	expectedSignature := h.Sum(nil)

	base64Signature := createBase64Encode(expectedSignature)

	if base64Signature != jwtSignature {
		http.Error(w, "Invalid JWT token signature", http.StatusUnauthorized)
		return
	}
		next.ServeHTTP(w, r)
	}
}

func createBase64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}