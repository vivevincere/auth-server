package authserver

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func TokenVerification(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["verify jwt claims here"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
