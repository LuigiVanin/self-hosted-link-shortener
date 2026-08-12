package link

type CreateLinkDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Url         string `json:"url" validate:"required,url"`
}

type FindLinkResponseDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Url         string `json:"url" validate:"required,url"`
	Code        string `json:"code"`
}

type FindLinkDto struct {
	Code string `json:"code"`
}
