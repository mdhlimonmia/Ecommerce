package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// JWT = JSON Web Token
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Payload is the data that we want to encode in the JWT
type Payload struct {
	Sub         string `json:"sub"` //user id call sub every where
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

// base64UrlEncode is a helper function that encodes data to base64 URL encoding without padding
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

// CreateJwt is a function that creates a JWT token from the given payload and secret key
func CreateJwt(data Payload, secret string) (string, error) {
	//create a header for the JWT token
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	//Convert the header to byte array and encode it to base64 URL encoding without padding
	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrayHeader)

	//Convert the payload to byte array and encode it to base64 URL encoding without padding
	byteArrayPayload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrayPayload)

	//The message that we want to sign is the concatenation of the header and payload separated by a dot
	message := headerB64 + "." + payloadB64
	byteArrayMessage := []byte(message)

	//Convert the secret key to byte array
	byteArraySecret := []byte(secret)

	//Create a new HMAC using sha256 hash function and the secret key
	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrayMessage)

	//Calculate the HMAC signature and encode it to base64 URL encoding without padding
	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	//The JWT token is the concatenation of the header, payload and signature separated by dots
	jwt := headerB64 + "." + payloadB64 + "." + signatureB64
	return jwt, nil
}
