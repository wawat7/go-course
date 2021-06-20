package exception

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"go-course/app/helper"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(validation.Errors)
	if ok {
		return ctx.JSON(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	
	return ctx.JSON(helper.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
