package middlewares

import (
	erro "selfhost-link-shortener/shared/errors"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	url := ctx.OriginalURL()

	if err, ok := err.(*erro.AppError); ok {
		return ctx.Status(err.Status).JSON(err.IntoProblemDetail(url))
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(
		erro.NewProblemDetail(
			"about:blank",
			"Unexpected Internal Error",
			fiber.StatusInternalServerError,
			err.Error(),
			url,
			"",
		),
	)
}
