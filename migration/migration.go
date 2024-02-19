package migration

import (
	"fmt"
	"gin-socmed/config"
	"gin-socmed/model"
	"log"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&model.User{},
		&model.Post{},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
