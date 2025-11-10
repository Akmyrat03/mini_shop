package main

import (
	"context"
	"log"

	"github.com/Akmyrat03/shop/internal/config"
	controller "github.com/Akmyrat03/shop/internal/controller/http/v0"
	"github.com/Akmyrat03/shop/internal/repository/postgres"
	"github.com/Akmyrat03/shop/internal/usecase"
	"github.com/Akmyrat03/shop/pkg/connection"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.App{}

	group := app.Group("/api/v0")

	cfg := config.LoadConfig()

	psqlDB, err := connection.NewDBConnection(context.Background(), cfg.Postgres)
	if err != nil {
		log.Printf("Failed to connect db: %v", err)
	}

	categoryRepo := postgres.NewCategoryRepository(psqlDB)

	categoryUC := usecase.NewCategoryUC(categoryRepo)

	categoryController := controller.NewCategoryUC(categoryUC)

	controller.MapCategoryRoutes(group, categoryController)
}
