package Controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"practice-bootcamp/Models"
	"practice-bootcamp/Utils"
)

func CreateBook(c *gin.Context) {
	var Input Models.InputBook

	err := c.ShouldBindJSON(&Input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	Book := Models.Book{Title: Input.Title, Author: Input.Author}

	Utils.DoCreate(&Book)

	c.JSON(http.StatusOK, gin.H{"Books": &Book})
}

func GetBooks(c *gin.Context) {
	var Books []Models.Book

	Utils.DoFind(&Books)

	c.JSON(http.StatusOK, gin.H{"Books": Books})
}

func FindBook(c *gin.Context) {
	var Book Models.Book

	err := Utils.IsPresent(c.Param("id"), &Book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Book Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Books": Book})
}

func UpdateBook(c *gin.Context) {
	var Book Models.Book

	err := Utils.IsPresent(c.Param("id"), &Book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Book id to be updated record not found"})
		return
	}

	var NewBook Models.UpdatedBook

	errNew := c.ShouldBindJSON(&NewBook)
	if errNew != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "New student request body is not correct"})
		return
	}

	Utils.DoUpdate(&Book, NewBook)

	c.JSON(http.StatusOK, NewBook)
}

func DeleteBook(c *gin.Context) {
	var Book Models.Book

	err := Utils.IsPresent(c.Param("id"), &Book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Book id to be deleted, record not found"})
		return
	}

	newErr := Utils.DoDelete(&Book)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occurred in deletion of the Book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Result": "Deleted Successfully"})
}
