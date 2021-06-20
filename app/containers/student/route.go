package student

import "github.com/gofiber/fiber/v2"

func (handler *Handler) Route(app *fiber.App)  {
	route := app.Group("api/")
	route.Get("students", handler.List)
}
