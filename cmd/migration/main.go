package main

import (
	"fmt"
	"selfhost-link-shortener/bootstrap"
	entity "selfhost-link-shortener/entities"
	"selfhost-link-shortener/shared"
)

func main() {
	env := shared.NewEnv()
	db := bootstrap.NewDatabase(env.DatabaseUrl)

	err := db.AutoMigrate(&entity.Link{})

	if err != nil {
		fmt.Println("Migration failed during constraint creation:", err.Error())
	}

	fmt.Println("Migrations Finished ✅")
}
