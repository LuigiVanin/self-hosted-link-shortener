package link

import (
	"fmt"
	"selfhost-link-shortener/shared"

	"github.com/LuigiVanin/openapi-builder/openapi"
	"github.com/gofiber/fiber/v3"
)

type LinkController struct {
	service LinkService
	swagger *openapi.Builder
}

func NewController(service LinkService, builder *openapi.Builder) LinkController {
	return LinkController{
		service: service,
		swagger: builder,
	}
}

func (this *LinkController) CreateShortLink(ctx fiber.Ctx) error {
	payload := CreateLinkDto{}
	err := ctx.Bind().Body(&payload)

	response, err := this.service.CreateShortLink(payload)

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	return ctx.Status(201).JSON(shared.JSON{
		"code": response.Code,
		"url":  fmt.Sprintf("%s/%s", ctx.BaseURL(), response.Code),
	})

}

func (this *LinkController) FindLink(ctx fiber.Ctx) error {
	code := ctx.Params("code")
	agent := ctx.UserAgent()

	link, err := this.service.FindLink(FindLinkDto{
		Code: code,
		Metadata: FindLinkMetadata{
			UserAgent: agent,
		},
	})

	if err != nil {
		return nil
	}

	if link.Body != "" {
		return ctx.Status(200).Send([]byte(link.Body))
	}

	return ctx.Redirect().To(link.Url)
}
