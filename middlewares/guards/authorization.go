package guards

import (
	b64 "encoding/base64"
	"fmt"
	"selfhost-link-shortener/shared"
	erro "selfhost-link-shortener/shared/errors"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func GetCredentials() map[string]any {
	env := shared.NewEnv()

	login := env.Login
	password := env.Password

	concat := fmt.Sprintf("%s%s", login, password)

	token := b64.StdEncoding.EncodeToString([]byte(concat))

	return shared.JSON{"token": token}
}

func Authorization(ctx fiber.Ctx) error {
	authorization := ctx.GetHeaders()["Authorization"]

	if len(authorization) == 0 {
		authorization = ctx.GetHeaders()["Auth"]
	}

	if (!ctx.HasHeader("Authorization") && !ctx.HasHeader("Auth")) || len(authorization) == 0 {
		return erro.ThrowBadRequest("Lack of authorization token")
	}

	authorization = strings.Split(authorization[0], " ")

	if len(authorization) < 2 {
		return erro.ThrowBadRequest("Malformatted token")
	}

	method := authorization[0]

	switch method {
	case "Basic":

		token := authorization[1]

		if token == GetCredentials()["token"] {
			return ctx.Next()
		}
	case "Bearer":
		return erro.ThrowNotImplementedError("Bearer token not implemented yet")
	}

	return erro.ThrowNotAllowed("Not Allowed use this route")
}
