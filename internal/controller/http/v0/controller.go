package controller

import (
	"log"

	"github.com/Akmyrat03/shop/internal/controller/http/v0/request"
	"github.com/Akmyrat03/shop/internal/controller/http/v0/response"
	"github.com/Akmyrat03/shop/internal/domain"
	"github.com/Akmyrat03/shop/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

var _ CategoryController = (*categoryController)(nil)

type CategoryController interface {
	CreateCategory(c *fiber.Ctx) error
}

type categoryController struct {
	categoryUC usecase.CategoryUC
}

func NewCategoryUC(categoryUC usecase.CategoryUC) *categoryController {
	return &categoryController{categoryUC: categoryUC}
}

func (h *categoryController) CreateCategory(c *fiber.Ctx) error {
	var req request.CreateCategoryReq
	err := c.BodyParser(&req)
	if err != nil {
		log.Printf("(c.BodyParser): %v", err)
		return c.SendStatus(400)
	}

	ctx := c.Context()

	category := domain.Category{
		NameTK: req.NameTK,
		NameEN: req.NameEN,
		NameRU: req.NameRU,
	}

	id, err := h.categoryUC.Create(ctx, category)
	if err != nil {
		log.Printf("(h.categoryUC.Create): %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	response := response.CreateCategoryRes{
		ID: id,
	}

	return c.Status(200).JSON(response)
}
