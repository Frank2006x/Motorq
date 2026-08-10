package router

import (
	"Frank2006xmotorq/internal/server/handler"

	"github.com/gofiber/fiber/v3"
)

func VehicleRoutes(app *fiber.App, h *handler.VehicleHandler) {
	vehicle := app.Group("/api/v1/vehicles")
	vehicle.Post("/", h.Create)
	vehicle.Get("/", h.ListByFleet)
	vehicle.Get("/:id", h.Get)
	vehicle.Put("/:id", h.Update)
	vehicle.Delete("/:id", h.Delete)
}