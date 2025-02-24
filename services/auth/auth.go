package auth

import (
	"fmt"
	"sinappsebackend/app"
	"sinappsebackend/services/users"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(app.JWTSecret), nil
}

func GenerateToken(u users.User) jwt.Token {
	return *jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": u.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
}

func IsValidToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return false
	}

	return token.Valid
}

func GetIdFromToken(tokenStr string) (uint32, error) {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := uint32(claims["id"].(float64))
		return id, nil
	}

	return 0, fmt.Errorf("cannot cast token claims to jwt.MapClaims")
}

func AuthenticateRequest(ctx *gin.Context) (bool, uint32, error) {
	authStr := ctx.GetHeader("Authorization")

	if !IsValidToken(authStr) {
		return false, 0, nil
	}

	id, err := GetIdFromToken(authStr)
	if err != nil {
		return false, 0, err
	}

	return true, id, nil
}