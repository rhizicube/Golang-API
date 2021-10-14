package put

import (
	"github.com/gin-gonic/gin"

	"log"

	"example/main/database"
)

func EditPeople(c *gin.Context) {
	id := c.Param("id")
	var newPerson database.People
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		log.Fatalf(err.Error())
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body"})
		return

	}
	for i, u := range database.Peoples {
		if u.ID == id {
			database.Peoples[i].Name = newPerson.Name
			database.Peoples[i].Contact = newPerson.Contact
			database.Peoples[i].City = newPerson.City

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
