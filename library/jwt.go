package library

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Admin bool `json:"admin"`
	Guid string
	jwt.StandardClaims
}

// 生成Jwt Token
func Jwt(guid string) (string, error) {
	claims := &jwtCustomClaims{
		true,
		guid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil

}

// 验证jwt token是否合法
func ParseToken(tokenString string) (*jwt.Token, *jwtCustomClaims, error) {
	claims := &jwtCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	return token, claims, err
}
