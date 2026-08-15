package link

import (
	"io"
	"net/http"
	"regexp"
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

func (this *LinkService) CreateShortLink(payload CreateLinkDto) (CreateLinkResponseDto, error) {
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
		return CreateLinkResponseDto{}, erro.ThrowInternalServerError("Unable to create short link")
	}

	return CreateLinkResponseDto{Code: code}, nil
}

func (this *LinkService) FindLink(payload FindLinkDto) (FindLinkResponseDto, error) {
	result := entity.Link{}

	err := this.client.Where(entity.Link{Code: payload.Code}).First(&result).Error

	if err != nil {
		return FindLinkResponseDto{}, erro.ThrowNotFound("Could not find the link")
	}

	if ok, _ := regexp.MatchString("(?i)whatsapp", payload.Metadata.UserAgent); ok {
		client := http.Client{
			Timeout: time.Second * 5,
		}

		response, err := client.Get(result.Url)

		if err == nil {
			defer response.Body.Close()
		}

		if err == nil && response.StatusCode == 200 {

			if raw, err := io.ReadAll(response.Body); err == nil {

				return FindLinkResponseDto{
					Body: string(raw),
				}, nil
			}
		}

	}

	return FindLinkResponseDto{
		Name:        result.Name,
		Url:         result.Url,
		Description: result.Description,
		Code:        result.Code,
	}, nil
}
