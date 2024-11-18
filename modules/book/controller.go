package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sb-go-quiz-week-3/database"
	"strconv"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	book, err := GetAllBookFrom(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book Book

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = InsertBookTo(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}

func GetBookById(c *gin.Context) {
	var book Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.Id = id

	var (
		result gin.H
	)

	book, err := GetBookByIdFrom(database.DbConnection, book)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteBook(c *gin.Context) {
	var book Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.Id = id
	err := DeleteBookFrom(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}
