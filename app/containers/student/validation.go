package student

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go-course/app/exception"
)

func Validate(request FormatRequest)  {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&request.Email, is.Email, validation.Required),
		validation.Field(&request.Gender, validation.Required),
		validation.Field(&request.PhoneNumber, validation.Length(1, 13)),
	)

	if err != nil {
		exception.PanicIfNeeded(err)
	}
}
