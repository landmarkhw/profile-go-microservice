package test

import (
	"time"

	"github.com/gin-gonic/gin"
)

func getJsonSimple(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func getJsonComplex(ctx *gin.Context) {
	var items [1000]gin.H
	var now = time.Now()

	// Generate 1000 items
	for i := 0; i < 1000; i++ {
		items[i] = gin.H{
			"id":      i,
			"another": "thing",
			"ts":      now,
		}
	}

	ctx.JSON(200, gin.H{
		"id":      1,
		"items":   items,
		"success": true,
	})
}

// Defines routes used by the web server
func Routes(engine *gin.Engine) {
	engine.GET("/test/json/simple", getJsonSimple)
	engine.GET("/test/json/complex", getJsonComplex)
}
