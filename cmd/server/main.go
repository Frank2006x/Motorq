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
	"github.com/gofiber/fiber/v3/middleware/logger"
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
	ruleEvaluator := service.NewRuleEvaluator(queries)
	simpleRuleFleetSvc := service.NewSimpleRuleFleetService(queries)
	simpleRuleVehicleSvc := service.NewSimpleRuleVehicleService(queries)

	fleetHandler := handler.NewFleetHandler(fleetSvc)
	vehicleHandler := handler.NewVehicleHandler(vehicleSvc)
	telemetryHandler := handler.NewTelemetryHandler(telemetrySvc, ruleEvaluator)
	simpleRuleFleetHandler := handler.NewSimpleRuleFleetHandler(simpleRuleFleetSvc)
	simpleRuleVehicleHandler := handler.NewSimpleRuleVehicleHandler(simpleRuleVehicleSvc)

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	router.FleetRoutes(app, fleetHandler)
	router.VehicleRoutes(app, vehicleHandler)
	router.TelemetryRoutes(app, telemetryHandler)
	router.SimpleRuleFleetRoutes(app, simpleRuleFleetHandler)
	router.SimpleRuleVehicleRoutes(app, simpleRuleVehicleHandler)

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