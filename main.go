package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todos struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {
	db, err := gorm.Open(postgres.Open("postgresql://neondb_owner:RPxL8QMq7lAI@ep-wispy-bar-a1kx78in.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"), &gorm.Config{})
	db.AutoMigrate(&Todos{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Database")
	}

	router := gin.Default()
	router.GET("/api/getTodos", func(c *gin.Context) {
		var allTodos Todos
		result := db.Find(&allTodos)
		fmt.Println(result.RowsAffected)
	})
	router.POST("/api/createTodo", func(c *gin.Context) {
		var newTodo Todos

		if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&newTodo)
	})
	router.Run(":3001")
}
