package db

import (
	"fmt"
	"log"

	"github.com/Zync1402/ZyncifyServer/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgresql://neondb_owner:RPxL8QMq7lAI@ep-wispy-bar-a1kx78in.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"), &gorm.Config{})
	db.AutoMigrate(&models.Todos{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Database")
	}

	return db
}
