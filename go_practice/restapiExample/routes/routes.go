package routes

import (
	"example.com/go_practice/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //events/1, events/5 etc //dynamic path handler.

	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events/register/:id", createRegistration)
	authenticated.DELETE("/events/register/:id", deleteRegistration)

}
