package shared

import "os"

type Env struct {
	Login    string
	Password string

	DatabaseUrl string
}

func NewEnv() Env {
	login := os.Getenv("USER_LOGIN")
	password := os.Getenv("USER_PASSWORD")
	databaseUrl := os.Getenv("DATABASE_URL")

	return Env{
		Login:       login,
		Password:    password,
		DatabaseUrl: databaseUrl,
	}
}
