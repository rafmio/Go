// https://bytegoblin.io/blog/open-api-with-go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.Run("localhost:8088")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}
