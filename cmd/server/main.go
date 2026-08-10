package main

import (
	"Frank2006xmotorq/internal/config"
	"Frank2006xmotorq/internal/db"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		os.Exit(1)
	}

	pool, err := db.NewDB(cfg.POSTGRES_DB)
	if err != nil {
		fmt.Printf("failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	fmt.Println("Connected to database!")

	app := fiber.New()
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Listen(":8080")
}