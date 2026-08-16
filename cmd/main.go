package main

import (
	"fmt"
	"selfhost-link-shortener/bootstrap"
	m "selfhost-link-shortener/middlewares"
	"selfhost-link-shortener/middlewares/guards"
	"selfhost-link-shortener/modules/link"
	"selfhost-link-shortener/shared"

	"github.com/LuigiVanin/openapi-builder/openapi"
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

	document := builder.Build()
	swagger := bootstrap.NewSwagger(*document)

	server.Get("/docs", swagger.Handler)
	server.Get("/docs/*", swagger.Handler)

	server.Get(
		"/:code",
		linkController.FindLink,
	)

	server.Listen(addr)

}
