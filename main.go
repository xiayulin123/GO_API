package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tool struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var tools = []tool{
	{ID: "1", Name: "Eraser", Price: 2, Quantity: 10},
	{ID: "2", Name: "Pen", Price: 4, Quantity: 9},
	{ID: "3", Name: "Pencil", Price: 3, Quantity: 6},
}

func getTools(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tools)
}
func toolById(c *gin.Context) {
	id := c.Param("id")
	tool, err := getToolById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no such id for tool"})
		return
	}
	c.IndentedJSON(http.StatusOK, tool)
}
func getToolById(id string) (*tool, error) {
	for i, k := range tools {
		if k.ID == id {
			return &tools[i], nil
		}
	}

	return nil, errors.New("No tool is found by id")
}

func newTools(c *gin.Context) {
	var newTool tool
	if err := c.BindJSON(&newTool); err != nil {
		return
	}

	tools = append(tools, newTool)
	c.IndentedJSON(http.StatusCreated, newTool)
}
func main() {
	router := gin.Default()
	router.GET("/tools", getTools)
	router.GET("/tools/:id", toolById)

	router.POST("/tools", newTools)
	router.Run("localhost:8080")

}
