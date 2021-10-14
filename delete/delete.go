package delete

import (
	"github.com/gin-gonic/gin"

	"log"

	"example/main/database"
)

func RemovePersonByID(c *gin.Context) {
	id := c.Param("id")

	for i, u := range database.Peoples {
		if u.ID == id {
			database.Peoples = append(database.Peoples[:i])

			c.JSON(200, gin.H{
				"error": false})
			return

		}
	}
	log.Println("Person not found")
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
