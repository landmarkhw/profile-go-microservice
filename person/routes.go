package person

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	person := Get(id)
	c.JSON(200, gin.H{"data": person})
}

func savePerson(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the person in the database
	Save(person)

	c.Status(http.StatusOK)
}

// Defines routes
func Routes(engine *gin.Engine) {
	engine.GET("/person/:id", getPerson)
	engine.PUT("/person", savePerson)
}
