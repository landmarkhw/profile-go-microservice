package status

import (
	"github.com/gin-gonic/gin"
	"github.com/landmarkhw/profile-go-microservice/database"
)

func getStatus(c *gin.Context) {
	dbSuccess := pfdb.Test()

	c.JSON(200, gin.H{
		"database": dbSuccess,
	})
}

// Defines routes used by the web server
func Routes(engine *gin.Engine) {
	engine.GET("/status", getStatus)
}
