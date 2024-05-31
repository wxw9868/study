package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signingKey = []byte("ledger")

type MyCustomClaims struct {
	Username string `json:"username,omitempty"`
	jwt.RegisteredClaims
}

// Create claims with multiple fields populated
var claims = MyCustomClaims{
	"bar",
	jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		Issuer:    "test",
		Subject:   "somebody",
		ID:        "1",
		Audience:  []string{"somebody_else"},
	},
}

// GenerateJwtToken 生成一个JWT令牌
func GenerateJwtToken() (string, error) {
	// 使用HS256签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// ParseJwtToken 函数解析JWT
// token并返回一个指向MyCustomClaims的指针和错误信息
// 如果解析成功，则返回一个指向MyCustomClaims的指针和一个nil错误信息
// 如果解析失败，则返回一个nil和一个非nil错误信息
func ParseJwtToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}
