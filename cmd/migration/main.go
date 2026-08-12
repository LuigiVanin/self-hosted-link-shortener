package main

import (
	"fmt"
	"selfhost-link-shortener/bootstrap"
	entity "selfhost-link-shortener/entities"
)

const URL string = "postgresql://postgres:1337@192.168.15.3:5432/selfhost-linkshortener"

func main() {
	db := bootstrap.NewDatabase(URL)

	err := db.AutoMigrate(&entity.Link{})

	if err != nil {
		fmt.Println("Migration failed during constraint creation:", err.Error())
	}

	fmt.Println("Migrations Finished ✅")
}
