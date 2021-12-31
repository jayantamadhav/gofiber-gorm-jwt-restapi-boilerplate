package utils

import (
	"github.com/gofiber/fiber/v2"
)

func Response(data interface{}, status_code int, success bool, message string) fiber.Map {
	return fiber.Map{
		"data":        data,
		"status_code": status_code,
		"success":     success,
		"message":     message,
	}
}
