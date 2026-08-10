package main

import (
	"fmt"
	"os"

	"Frank2006xmotorq/db/sqlc"
	"Frank2006xmotorq/internal/config"
	"Frank2006xmotorq/internal/db"
	"Frank2006xmotorq/internal/server/handler"
	"Frank2006xmotorq/internal/server/router"
	"Frank2006xmotorq/internal/service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.MustLoadConfig(".")

	pool, err := db.NewDB(cfg.POSTGRES_DB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	fmt.Println("Connected to database!")

	queries := sqlc.New(pool)

	fleetSvc := service.NewFleetService(queries)
	vehicleSvc := service.NewVehicleService(queries)
	telemetrySvc := service.NewTelemetryService(queries)

	fleetHandler := handler.NewFleetHandler(fleetSvc)
	vehicleHandler := handler.NewVehicleHandler(vehicleSvc)
	telemetryHandler := handler.NewTelemetryHandler(telemetrySvc)

	app := fiber.New()

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	router.FleetRoutes(app, fleetHandler)
	router.VehicleRoutes(app, vehicleHandler)
	router.TelemetryRoutes(app, telemetryHandler)

	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on :%s\n", port)
	if err := app.Listen(":" + port); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start server: %v\n", err)
		os.Exit(1)
	}
}