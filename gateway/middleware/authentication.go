package middleware

import "github.com/gofiber/fiber"

func Authentication(c *fiber.Ctx) {
	c.Next()
}
