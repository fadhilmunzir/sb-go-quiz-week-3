package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sb-go-quiz-week-3/database"
	"strconv"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := GetAllCategoryFrom(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category Category

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = InsertCategoryTo(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

func GetCategoryById(c *gin.Context) {
	var category Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.Id = id

	var (
		result gin.H
	)

	category, err := GetCategoryByIdFrom(database.DbConnection, category)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteCategory(c *gin.Context) {
	var category Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.Id = id
	err := DeleteCategoryFrom(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

func GetAllBookInCategory(c *gin.Context) {
	var category Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.Id = id
	var (
		result gin.H
	)

	book, err := GetAllBookInCategoryFrom(database.DbConnection, category)

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
