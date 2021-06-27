package router

import (
	"github.com/gofiber/fiber/v2"
)

// Router is router of api
func (app *App) Router(r *fiber.App) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	email := emailVerifierHandler{}
	v1.Get("/verify", email.verifyHandler)
}
