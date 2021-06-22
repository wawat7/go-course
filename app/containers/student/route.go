package student

import "github.com/gofiber/fiber/v2"

func (handler *Handler) Route(route fiber.Router)  {
	route.Get("students", handler.List)
	route.Post("students", handler.Create)
	route.Get("students/:id", handler.FindById)
}
