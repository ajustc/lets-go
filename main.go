package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Restlaith struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var restlaith = []Restlaith{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Learn Docker", Completed: false},
	{ID: "3", Item: "Learn Kubernetes", Completed: false},
	{ID: "4", Item: "Learn AWS", Completed: false},
	{ID: "5", Item: "Learn Azure", Completed: false},
	{ID: "6", Item: "Learn GCP", Completed: false},
	{ID: "7", Item: "Learn DevOps", Completed: false},
	{ID: "8", Item: "Learn CICD", Completed: false},
	{ID: "9", Item: "Learn DevSecOps", Completed: false},
	{ID: "10", Item: "Learn Terraform", Completed: false},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, restlaith)
}

func addItem(c *gin.Context) {
	var newItem Restlaith

	if err := c.BindJSON(&newItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": strings.Replace(err.Error(), "json: ", "", -1),
		})
		return
	}

	restlaith = append(restlaith, newItem)

	c.IndentedJSON(http.StatusCreated, restlaith)
}

func getItemById(id string) (*Restlaith, error) {
	for i, t := range restlaith {
		if t.ID == id {
			return &restlaith[i], nil
		}
	}

	return nil, errors.New("item not found")
}

func getItem(c *gin.Context) {
	id := c.Param("id")

	item, err := getItemById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.GET("/items/:id", getItem)
	// router.PUT("/items/:id", updateItem)
	// router.DELETE("/items/:id", deleteItem)

	fmt.Fprintln(os.Stderr, "[Justc-Dev] Development running on localhost:2020")
	router.Run("localhost:2020")
}
