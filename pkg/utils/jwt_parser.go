package utils

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func ExtractTokenMetadata(c fiber.Ctx) (*jwt.RegisteredClaims, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func extractToken(c fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func verifyToken(c fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(viper.GetString("jwt.secret-key")), nil
}
