package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type people struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	City    string `json:"city"`
}

var peoples = []people{
	{ID: "1", Name: "Sonia", Contact: "sonia@gmail.com", City: "Bangalore"},
	{ID: "2", Name: "Shubhi", Contact: "shubhi@gmail.com", City: "Raipur"},
	{ID: "3", Name: "Lalit", Contact: "lalit@gmail.com", City: "Delhi"},
	{ID: "4", Name: "Udit", Contact: "udit@gmail.com", City: "Singapore"},
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, peoples)
}

func postPeople(c *gin.Context) {
	var newPerson people
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body"})
		return
	}
	peoples = append(peoples, newPerson)
	c.JSON(200, gin.H{
		"error": false})

}
func editPeople(c *gin.Context) {
	id := c.Param("id")
	var newPerson people
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body"})
		return

	}
	for i, u := range peoples {
		if u.ID == id {
			peoples[i].Name = newPerson.Name
			peoples[i].Contact = newPerson.Contact
			peoples[i].City = newPerson.City

			c.JSON(200, gin.H{
				"error": false})
			return

		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range peoples {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}
func removePersonByID(c *gin.Context) {
	id := c.Param("id")

	for i, u := range peoples {
		if u.ID == id {
			peoples = append(peoples[:i])

			c.JSON(200, gin.H{
				"error": false})
			return

		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people", postPeople)
	router.PUT("/people/:id", editPeople)
	router.GET("/people/:id", getPersonByID)
	router.DELETE("/people/:id", removePersonByID)

	router.Run("localhost:8000")
}
