package bootstrap

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(databaseUrl string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		fmt.Println("Erro ao rodar gorm open")
		panic(err.Error())
	}

	return db
}
