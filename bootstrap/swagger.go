package bootstrap

import (
	"net/http"

	"github.com/flowchartsman/swaggerui"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"

	"github.com/LuigiVanin/openapi-builder/openapi"
)

type Swagger struct {
	Handler fiber.Handler
	adapter fiber.Handler
}

func NewSwagger(document openapi.Document) Swagger {
	content, err := document.Output("yaml")

	if err != nil {
		panic(err.Error())
	}

	swagger := Swagger{
		// Handler: handler,
		adapter: adaptor.HTTPHandler(
			http.StripPrefix("/docs", swaggerui.Handler([]byte(content))),
		),
	}

	swagger.Handler = func(ctx fiber.Ctx) error {
		if ctx.Path() == "/docs" {
			return ctx.Redirect().To("/docs/")
		}

		return swagger.adapter(ctx)
	}

	return swagger
}
