package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// base64UrlEncode is a helper function that encodes data to base64 URL encoding without padding
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("############################AuthenticateJWT middleware called")
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorize", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]
		accessTokenArr := strings.Split(accessToken, ".")
		if len(accessTokenArr) != 3 {
			http.Error(w, "Unauthorize", http.StatusUnauthorized)
			return
		}

		jwtHeader := accessTokenArr[0]
		jwtPayload := accessTokenArr[1]
		jwtSignature := accessTokenArr[2]

		message := jwtHeader + "." + jwtPayload

		cnf := m.config

		byteArrSecret := []byte(cnf.JwtSecretKey)
		byteArrMessage := []byte(message)
		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		//Calculate the HMAC signature and encode it to base64 URL encoding without padding
		signature := h.Sum(nil)
		signatureB64 := base64UrlEncode(signature)

		if jwtSignature != signatureB64 {
			http.Error(w, "Unauthorize", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
