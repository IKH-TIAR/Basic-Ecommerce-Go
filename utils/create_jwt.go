package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) string {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArryHeader, err := json.Marshal(header)
	if err!=nil {
		return err.Error()
	}

	headerB64 := base64URLEncode(byteArryHeader)

	byteArryPayload, err := json.Marshal(data)
	if err!=nil {
		return err.Error()
	}
	payloadB64 := base64URLEncode(byteArryPayload)

	meassage := headerB64 + "." + payloadB64

	byteArrySecret := []byte(secret)
	byteArryMessage := []byte(meassage)

	h := hmac.New(sha256.New, byteArrySecret)
	h.Write(byteArryMessage)

	signature := h.Sum(nil)

	signatureB64 := base64URLEncode(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt

}

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}