package link

type CreateLinkDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Url         string `json:"url" validate:"required,url"`
}

type CreateLinkResponseDto struct {
	// Url  string `json:"url"`
	Code string `json:"code"`
}

type FindLinkResponseDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Url         string `json:"url" validate:"required,url"`
	Code        string `json:"code"`

	Body string `json:"-"`
}

type FindLinkMetadata struct {
	UserAgent string `json:"user-agent"`
}

type FindLinkDto struct {
	Code     string           `json:"code"`
	Metadata FindLinkMetadata `json:"metadata"`
}
