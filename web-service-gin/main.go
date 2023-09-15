package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "56", Title: "Fahrenheıt 451", Artist: "Ray Bradbury", Price: 56.99},
	{ID: "27", Title: "Yüksek Ökçeler", Artist: "Ömer Seyfettin", Price: 17.99},
	{ID: "3", Title: "Sıfırdan Zirveye", Artist: "Jeff Bezos", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "kitap bulunamadı"})
}
