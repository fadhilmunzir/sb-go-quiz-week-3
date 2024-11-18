package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"sb-go-quiz-week-3/database"
	"sb-go-quiz-week-3/modules/book"
	"sb-go-quiz-week-3/modules/category"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()
	router.GET("/books", book.GetAllBook)
	router.POST("/books", book.InsertBook)
	router.GET("/books/:id", book.GetBookById)
	router.DELETE("/books/:id", book.DeleteBook)

	router.GET("/categories", category.GetAllCategory)
	router.POST("/categories", category.InsertCategory)
	router.GET("/categories/:id", category.GetCategoryById)
	router.DELETE("/categories/:id", category.DeleteCategory)

	router.Run(":8080")
}
