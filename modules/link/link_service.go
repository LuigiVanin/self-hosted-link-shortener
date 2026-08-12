package link

import (
	"fmt"
	entity "selfhost-link-shortener/entities"
	"selfhost-link-shortener/shared"
	erro "selfhost-link-shortener/shared/errors"
	"time"

	"gorm.io/gorm"
)

type LinkService struct {
	client *gorm.DB
}

func NewService(client *gorm.DB) LinkService {
	return LinkService{
		client: client,
	}
}

func (this *LinkService) CreateShortLink(payload CreateLinkDto) (string, error) {
	code := shared.GenerateRandomCharacters(6)

	link := entity.Link{
		Name:        payload.Name,
		Code:        code,
		Description: payload.Description,
		Url:         payload.Url,
		CreatedAt:   time.Now(),
	}

	err := this.client.Create(&link).Error

	if err != nil {
		fmt.Println(err.Error())
		return "", erro.ThrowInternalServerError("Unable to create short link")
	}

	return code, nil
}

func (this *LinkService) FindLink(payload FindLinkDto) (FindLinkResponseDto, error) {
	result := entity.Link{}

	err := this.client.Where(entity.Link{Code: payload.Code}).First(&result).Error

	if err != nil {
		return FindLinkResponseDto{}, erro.ThrowNotFound("Could not find the link")
	}

	return FindLinkResponseDto{
		Name:        result.Name,
		Url:         result.Url,
		Description: result.Description,
		Code:        result.Code,
	}, nil
}
