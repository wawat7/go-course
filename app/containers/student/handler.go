package student

import (
	"github.com/gofiber/fiber/v2"
	"go-course/app/helper"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service2 Service) *Handler {
	return &Handler{service: service2}
}

func (handler *Handler) List(c *fiber.Ctx) error {
	responses := handler.service.List()
	return c.JSON(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "List Students",
		Data:   responses,
	})
}