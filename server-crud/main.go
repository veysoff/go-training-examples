package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// albums slice to seed record album data.
var tasks = []task{
	{ID: "1", Title: "Title 1"},
	{ID: "2", Title: "Title 2"},
	{ID: "3", Title: "Title 3"},
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// postAlbums adds an album from JSON received in the request body.
func postTasks(c *gin.Context) {
	var newTask task

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}
