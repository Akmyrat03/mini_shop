package controller

import "github.com/gofiber/fiber/v2"

func MapCategoryRoutes(r fiber.Router, h CategoryController) {
	r.Post("/category", h.CreateCategory)
}
