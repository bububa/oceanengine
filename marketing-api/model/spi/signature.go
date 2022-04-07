package spi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sign 生成签名
func Sign(secret []byte, data []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// IsValidToken 验证签名
func IsValidToken(secret []byte, token []byte, data []byte) bool {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	expect := h.Sum(nil)
	return hmac.Equal(token, expect)
}

// VerifyTokenHeader 验证签名
func VerifyTokenHeader(secret []byte, data []byte, tokenFromHeader string) bool {
	token, err := hex.DecodeString(tokenFromHeader)
	if err != nil {
		return false
	}
	return IsValidToken(secret, token, data)
}
