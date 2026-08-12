package bootstrap

import (
	"selfhost-link-shortener/middlewares"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Json(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	return ctx.Next()
}

func NewHttpServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Link Shortener",
		ServerHeader: "Link Shortener",
		ErrorHandler: middlewares.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	app.Use(Json)

	return app
}
