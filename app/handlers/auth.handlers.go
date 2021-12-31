package handlers

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewToken() (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
func GetToken(c *fiber.Ctx) error {
	token, err := GenerateNewToken()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Couldn't generate a token",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Request successful",
		"data": fiber.Map{
			"access_token": &token,
		},
	})
}
