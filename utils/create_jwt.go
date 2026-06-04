package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg  string `json:"alg"`
	Type string `json:"typ"`
}

type Payload struct {
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"userType"`
	Iat      int64  `json:"iat,omitempty"`
	Exp      int64  `json:"exp,omitempty"`
}

type Signature struct {
	Signature string `json:"signature"`
}

func CreateJWT( payload Payload) (string, error) {
	// In a real implementation, you would use a library to create a JWT token
	// Here we are just simulating the process for demonstration purposes

	headers := Header{
		Alg:  "HS256",
		Type: "JWT",
	}

	bytesHeader, err := json.Marshal(headers)
	if err != nil {
		panic(err)
	}

	headerBase64 := createBase64Encode(bytesHeader)

	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	payloadBase64 := createBase64Encode(bytesPayload)

	message := headerBase64 + "." + payloadBase64 

	secret := "my_secret_key"
	signatureByteArray := []byte(secret)
	bytesMessage := []byte(message)

	hMac := hmac.New(sha256.New, signatureByteArray)
	hMac.Write(bytesMessage)

	signature := hMac.Sum(nil)

	signatureBase64 := createBase64Encode(signature)

	jwt := message + "." + signatureBase64

	return jwt, nil
	

}

func createBase64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
