package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/landmarkhw/profile-go-microservice/person"
	"github.com/landmarkhw/profile-go-microservice/status"
	"github.com/landmarkhw/profile-go-microservice/test"
)

// Run the REST API web server
func Run() {
	engine := gin.Default()

	person.Routes(engine)
	status.Routes(engine)
	test.Routes(engine)

	engine.Run(":5000")
}
