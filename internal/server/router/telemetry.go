package router

import (
	"Frank2006xmotorq/internal/server/handler"

	"github.com/gofiber/fiber/v3"
)

func TelemetryRoutes(app *fiber.App, h *handler.TelemetryHandler) {
	telemetry := app.Group("/api/v1/telemetry")
	telemetry.Post("/", h.Create)
	telemetry.Get("/:id", h.Get)
	telemetry.Get("/vehicle", h.ListByVehicle)
	telemetry.Get("/vehicle/latest", h.ListLatestByVehicle)
}