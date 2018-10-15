package test

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CustomType struct {
	Prop1 int
	Prop2 string
	Prop3 time.Time
}

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

func getJsonCustom(ctx *gin.Context) {
	custom := CustomType{
		Prop1: 1,
		Prop2: "testing",
		Prop3: time.Now(),
	}
	ctx.JSON(200, custom)
}

func getHelloWorld(ctx *gin.Context) {
	ctx.Writer.WriteString("Hello, world!")
	ctx.Status(200)	
}

// Defines routes used by the web server
func Routes(engine *gin.Engine) {
	engine.GET("/test/json/simple", getJsonSimple)
	engine.GET("/test/json/complex", getJsonComplex)
	engine.GET("/test/json/custom", getJsonCustom)
	engine.GET("/test/json/hello-world", getHelloWorld)
}
