package get

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"example/main/database"
)

func GetPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Peoples)
}

func GetPersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range database.Peoples {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	log.Println("Person not found")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}
