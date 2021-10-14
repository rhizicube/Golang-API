package post

import (
	"github.com/gin-gonic/gin"

	"log"

	"example/main/database"
)

func PostPeople(c *gin.Context) {
	var newPerson database.People
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		log.Fatalf(err.Error())

		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body"})
		return
	}

	database.Peoples = append(database.Peoples, newPerson)
	c.JSON(200, gin.H{
		"error": false})

}
