package main

import (
	"example/main/delete"

	"github.com/gin-gonic/gin"

	"example/main/get"

	"example/main/put"

	"example/main/post"
)

func main() {
	router := gin.Default()
	router.GET("/people", get.GetPeople)
	router.POST("/people", post.PostPeople)
	router.PUT("/people/:id", put.EditPeople)
	router.GET("/people/:id", get.GetPersonByID)
	router.DELETE("/people/:id", delete.RemovePersonByID)

	router.Run("localhost:8000")
}
