package router

import (
	"Frank2006xmotorq/internal/server/handler"

	"github.com/gofiber/fiber/v3"
)

func SimpleRuleVehicleRoutes(app *fiber.App, h *handler.SimpleRuleVehicleHandler) {
	rules := app.Group("/api/v1/simple-rule-vehicle")
	rules.Post("/", h.Create)
	rules.Get("/:id", h.Get)
	rules.Get("/", h.ListByVehicle)
	rules.Delete("/:id", h.Delete)
}