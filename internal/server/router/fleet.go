package router

import (
	"Frank2006xmotorq/internal/server/handler"

	"github.com/gofiber/fiber/v3"
)

func FleetRoutes(app *fiber.App, h *handler.FleetHandler) {
	fleet := app.Group("/api/v1/fleets")
	fleet.Post("/", h.Create)
	fleet.Get("/", h.List)
	fleet.Get("/:id", h.Get)
	fleet.Put("/:id", h.Update)
	fleet.Delete("/:id", h.Delete)
}