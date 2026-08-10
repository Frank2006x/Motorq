package router

import (
	"Frank2006xmotorq/internal/server/handler"

	"github.com/gofiber/fiber/v3"
)

func SimpleRuleFleetRoutes(app *fiber.App, h *handler.SimpleRuleFleetHandler) {
	rules := app.Group("/api/v1/simple-rule-fleet")
	rules.Post("/", h.Create)
	rules.Get("/:id", h.Get)
	rules.Get("/", h.ListByFleet)
	rules.Delete("/:id", h.Delete)
}