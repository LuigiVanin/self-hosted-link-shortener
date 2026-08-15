package main

import (
	"fmt"
	"net/http"
	"selfhost-link-shortener/bootstrap"
	m "selfhost-link-shortener/middlewares"
	"selfhost-link-shortener/middlewares/guards"
	"selfhost-link-shortener/modules/link"
	"selfhost-link-shortener/shared"

	"github.com/LuigiVanin/openapi-builder/openapi"
	"github.com/flowchartsman/swaggerui"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func main() {
	fmt.Println("Starting Server!")

	env := shared.NewEnv()
	client := bootstrap.NewDatabase(env.DatabaseUrl)
	server := bootstrap.NewHttpServer()
	builder := openapi.NewBuilder(
		"Link Shortener API",
		"A very simple self hosted link shortener",
		"1.0.0",
	)

	addr := fmt.Sprintf("0.0.0.0:%s", "8080")

	linkService := link.NewService(client)

	linkController := link.NewController(linkService, builder)

	builder.Add(
		builder.Route("POST", "/").
			AddBody(link.CreateLinkDto{}, openapi.Options{Required: true}).
			AddResponse(201, shared.JSON{}).
			AddHeaderParam("Auth", "string", openapi.Options{Required: true}),
	)
	server.Post(
		"/",
		guards.Authorization,
		m.Validator[link.CreateLinkDto](),
		linkController.CreateShortLink,
	)

	// path param sempre required: true -- a spec OpenAPI exige esse valor
	// para parametros em "path".
	builder.Add(
		builder.Route("GET", "/{code}").
			AddPathParam("code", "string", openapi.Options{Required: true}),
	)

	document := builder.Build()
	content, err := document.Output("json")

	if err != nil {
		panic(err.Error())
	}

	content, err = bootstrap.WithSecurityScheme(content, map[string][]string{
		"/": {"post"},
	})

	if err != nil {
		panic(err.Error())
	}

	// swaggerui.Handler devolve um http.ServeMux que espera ser montado na raiz,
	// entao o prefixo /docs precisa ser removido antes de chegar nele.
	swagger := adaptor.HTTPHandler(
		http.StripPrefix("/docs", swaggerui.Handler([]byte(content))),
	)

	docsHandler := func(ctx fiber.Ctx) error {
		if ctx.Path() == "/docs" {
			return ctx.Redirect().To("/docs/")
		}

		return swagger(ctx)
	}

	server.Get("/docs", docsHandler)
	server.Get("/docs/*", docsHandler)

	server.Get(
		"/:code",
		linkController.FindLink,
	)

	server.Listen(addr)

}
