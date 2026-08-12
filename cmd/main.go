package main

import (
	"fmt"
	"selfhost-link-shortener/bootstrap"
	m "selfhost-link-shortener/middlewares"
	"selfhost-link-shortener/middlewares/guards"
	"selfhost-link-shortener/modules/link"
	"selfhost-link-shortener/shared"
)

func main() {
	fmt.Println("Starting Server!")

	env := shared.NewEnv()
	client := bootstrap.NewDatabase(env.DatabaseUrl)
	server := bootstrap.NewHttpServer()

	addr := fmt.Sprintf("0.0.0.0:%s", "8080")

	linkService := link.NewService(client)

	linkController := link.NewController(linkService)

	server.Post(
		"/",
		guards.Authorization,
		m.Validator[link.CreateLinkDto](),
		linkController.CreateShortLink,
	)

	server.Get(
		"/:code",
		linkController.FindLink,
	)

	server.Listen(addr)

}
