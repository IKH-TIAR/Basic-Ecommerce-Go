package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/utils"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Authenticate Middleware")

		header := r.Header.Get("Authorization")

		if header == "" {
			utils.WriteError(w, http.StatusUnauthorized, "Unauthorized 1")
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			utils.WriteError(w, http.StatusUnauthorized, "Unauthorized 2")
			return
		}

		if headerArr[0] != "Bearer" {
			utils.WriteError(w, http.StatusUnauthorized, "Unauthorized 3")
			return
		}

		tokenParts := strings.Split(headerArr[1], ".")

		if len(tokenParts) != 3 {
			utils.WriteError(w, http.StatusUnauthorized, "Unauthorized 4")
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		meassage := jwtHeader + "." + jwtPayload

		byteArrySecret := []byte(config.GetConfig().Secret)
		byteArryMessage := []byte(meassage)

		h := hmac.New(sha256.New, byteArrySecret)
		h.Write(byteArryMessage)

		signature := h.Sum(nil)

		newSignatureB64 := base64URLEncode(signature)

		if newSignatureB64 != jwtSignature {
			utils.WriteError(w, http.StatusUnauthorized, "Hacker beda tui")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
