package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var tools = []book{
	{ID: "1", Name: "Eraser", Price: 2, Quantity: 10},
	{ID: "2", Name: "Pen", Price: 4, Quantity: 9},
	{ID: "3", Name: "Pencil", Price: 3, Quantity: 6},
}

func getTools(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tools)
}

func main() {
	router := gin.Default()
	router.GET("/books", getTools)
	router.Run("localhost:8080")

}
