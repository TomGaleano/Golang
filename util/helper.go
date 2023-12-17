package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))

}

func ParseJwt(token string) (string, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Issuer, err
}
