package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/mailchecker/app/email/services"
	"github.com/muhfaris/mailchecker/gateway/structures"
)

type emailVerifierHandler struct {
	service *services.EmailVerifier
}

func (h *emailVerifierHandler) verifyHandler(c *fiber.Ctx) error {
	var email structures.EmailVerifierRead
	if err := c.QueryParser(&email); err != nil {
		return c.JSON(Response{Status: http.StatusBadRequest, Error: err})
	}

	data, err := h.service.Validate(c.Context(), email)
	if err != nil {
		return c.JSON(Response{Status: http.StatusBadRequest, Mesage: err.Error()})
	}

	return c.JSON(Response{Status: http.StatusOK, Data: data})
}
