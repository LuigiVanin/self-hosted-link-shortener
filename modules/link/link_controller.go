package link

import (
	"fmt"
	"selfhost-link-shortener/shared"

	"github.com/gofiber/fiber/v3"
)

type LinkController struct {
	service LinkService
}

func NewController(service LinkService) LinkController {
	return LinkController{
		service: service,
	}
}

func (this *LinkController) CreateShortLink(ctx fiber.Ctx) error {
	payload := CreateLinkDto{}
	err := ctx.Bind().Body(&payload)

	fmt.Println("Payload: ", payload)

	code, err := this.service.CreateShortLink(payload)

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	return ctx.Status(201).JSON(shared.JSON{
		"code": code,
	})

}

func (this *LinkController) FindLink(ctx fiber.Ctx) error {
	code := ctx.Params("code")

	fmt.Println("Code: ", code)

	link, err := this.service.FindLink(FindLinkDto{
		Code: code,
	})

	if err != nil {
		return nil
	}

	return ctx.Redirect().To(link.Url)
}
