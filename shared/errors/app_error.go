package erro

import (
	"fmt"
	"maps"
	"selfhost-link-shortener/shared"

	"github.com/gofiber/fiber/v3"
)

type AppErrorCode string

type AppError struct {
	error
	Title  string
	Code   ErrorCodePair
	Status int
	Detail string
	Type   string
	Extra  shared.JSON
}

func NewAppError(title string, detail string, code ErrorCodePair, type_ string, extra ...shared.JSON) *AppError {

	extraInfo := make(shared.JSON)

	for _, extra := range extra {
		maps.Copy(extraInfo, extra)
	}

	return &AppError{
		Title:  title,
		Code:   code,
		Status: code.Second,
		Detail: detail,
		Extra:  extraInfo,
		Type:   type_,
	}
}

func ThrowNotAllowed(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Not Allowed",
		detail,
		NotAllowedErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/405",
		extra...,
	)
}

func ThrowTooManyRequests(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Too Many Requests",
		detail,
		TooManyRequestsErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/429",
		extra...,
	)
}

func ThrowBadRequest(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Bad Request",
		detail,
		BadRequestCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/400",
		extra...,
	)
}

func ThrowValidationError(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Bad Request",
		detail,
		ValidationErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/400",
		extra...,
	)
}

func ThrowInvalidOtpCode(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Invalid OTP Code",
		detail,
		InvalidOtpCodeErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/400",
		extra...,
	)
}

func ThrowConflict(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Conflict",
		detail,
		ConflictErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/409",
		extra...,
	)
}

func ThrowNotFound(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Not Found",
		detail,
		NotFoundErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/404",
		extra...,
	)
}

func ThrowUnauthorizedError(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Unauthorized",
		detail,
		UnauthorizedErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/401",
		extra...,
	)
}

func ThrowUnprocessableEntity(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Unprocessable Entity",
		detail,
		UnprocessableEntityErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/422",
		extra...,
	)
}

func ThrowInternalServerError(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Internal Server Error",
		detail,
		InternalServerErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/500",
		extra...,
	)
}

func ThrowNotImplementedError(detail string, extra ...shared.JSON) *AppError {
	return NewAppError(
		"Not Implemented",
		detail,
		NotImplementedErrorCode,
		"https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Reference/Status/501",
		extra...,
	)
}

func (e *AppError) Error() string {

	return fmt.Sprintf("AppError: %s, Code: %s, Detail: %s", e.Title, e.Code.First, e.Detail)
}

func (e *AppError) IntoProblemDetail(instance string) *ProblemDetail {
	status := e.Code.Second

	if status == 0 {
		status = fiber.StatusInternalServerError
	}

	return &ProblemDetail{
		Type:     e.Type,
		Title:    e.Title,
		Detail:   e.Detail,
		Instance: instance,
		Code:     string(e.Code.First),
		Status:   status,
		Data:     e.Extra,
	}
}
