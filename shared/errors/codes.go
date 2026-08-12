package erro

import (
	"selfhost-link-shortener/shared"

	"github.com/gofiber/fiber/v3"
)

type ErrorCodePair = shared.Pair[AppErrorCode, int]

var (
	// Internal Server 500
	InternalServerErrorCode ErrorCodePair = ErrorCodePair{
		First:  "INTERNAL_SERVER_ERROR",
		Second: fiber.StatusInternalServerError,
	}

	// Not Implemented 501
	NotImplementedErrorCode ErrorCodePair = ErrorCodePair{
		First:  "NOT_IMPLEMENTED",
		Second: fiber.StatusNotImplemented,
	}

	// Bad Request 400
	BadRequestCode ErrorCodePair = ErrorCodePair{
		First:  "BAD_REQUEST",
		Second: fiber.StatusBadRequest,
	}

	ValidationErrorCode ErrorCodePair = ErrorCodePair{
		First:  "VALIDATION_ERROR",
		Second: fiber.StatusBadRequest,
	}

	InvalidOtpCodeErrorCode ErrorCodePair = ErrorCodePair{
		First:  "INVALID_OTP_CODE",
		Second: fiber.StatusBadRequest,
	}

	// Unauthorized 401
	UnauthorizedErrorCode ErrorCodePair = ErrorCodePair{
		First:  "UNAUTHORIZED",
		Second: fiber.StatusUnauthorized,
	}

	// Not Found 404
	NotFoundErrorCode ErrorCodePair = ErrorCodePair{
		First:  "NOT_FOUND",
		Second: fiber.StatusNotFound,
	}

	// Method Not Allowed 405
	NotAllowedErrorCode ErrorCodePair = ErrorCodePair{
		First:  "NOT_ALLOWED",
		Second: fiber.StatusMethodNotAllowed,
	}

	// Too Many Requests 429
	TooManyRequestsErrorCode ErrorCodePair = ErrorCodePair{
		First:  "TOO_MANY_REQUESTS",
		Second: fiber.StatusTooManyRequests,
	}

	// Conflict 409
	ConflictErrorCode ErrorCodePair = ErrorCodePair{
		First:  "CONFLICT",
		Second: fiber.StatusConflict,
	}
	// Unprocessable Entity 422
	UnprocessableEntityErrorCode ErrorCodePair = ErrorCodePair{
		First:  "UNPROCESSABLE_ENTITY",
		Second: fiber.StatusUnprocessableEntity,
	}
)
