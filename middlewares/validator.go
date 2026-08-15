package middlewares

import (
	"fmt"
	"selfhost-link-shortener/shared"
	erro "selfhost-link-shortener/shared/errors"

	v "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type ValidationFieldError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Param   string `json:"param"`
	Value   any    `json:"value"`
	Details string `json:"details"`
}

type ValidationError struct {
	error
	List []ValidationFieldError `json:"list"`
}

func Validate[T any](payload T) error {
	validator := v.New()

	err := validator.Struct(payload)

	if err != nil {

		if _, ok := err.(*v.InvalidValidationError); ok {
			return erro.ThrowUnprocessableEntity(err.Error())
		}

		if err, ok := err.(v.ValidationErrors); ok {

			fields := []ValidationFieldError{}

			for _, e := range err {
				fmt.Println(e.Error())
				fields = append(
					fields,
					ValidationFieldError{
						Field:   e.Field(),
						Tag:     e.Tag(),
						Param:   e.Param(),
						Value:   e.Value(),
						Details: e.Error(),
					})
			}

			fmt.Println(fields)

			return erro.ThrowValidationError(err.Error(), shared.JSON{"fields": fields})
		}

		return erro.ThrowValidationError(err.Error())
	}

	return nil
}

func Validator[T any]() func(fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {

		var payload T
		err := ctx.Bind().Body(&payload)

		if err != nil {
			fmt.Println(err.Error())
			return erro.ThrowUnprocessableEntity(err.Error())
		}

		err = Validate(payload)

		if err != nil {

			return err
		}

		return ctx.Next()
	}
}
