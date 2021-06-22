package student

import (
	"github.com/gofiber/fiber/v2"
	"go-course/app/exception"
	"go-course/app/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "List Students",
		Data:   responses,
	})
}

func (handler *Handler) Create(c *fiber.Ctx) error {
	var request FormatRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	
	response := handler.service.Create(request)
	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Data has been saved",
		Data:   response,
	})
}

func (handler *Handler) FindById(c *fiber.Ctx) error {
	ID, err := primitive.ObjectIDFromHex(c.Params("id"))
	exception.PanicIfNeeded(err)
	response := handler.service.FindById(ID)
	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Data Student",
		Data:   response,
	})
}