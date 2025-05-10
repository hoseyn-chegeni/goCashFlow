package api

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes handles GET /hi
// @Summary SetupRoutes Endpoint
// @Description Responds with a simple "hi!" message.
// @Tags SetupRoutes
// @Accept json
// @Produce json
// @Success 200 {string} string "SetupRoutes!"
// @Router /hi/ [get]
func SetupRoutes(c *fiber.Ctx) error {
	return c.SendString("hi")
}
