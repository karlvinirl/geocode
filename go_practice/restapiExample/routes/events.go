package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/go_practice/restapi/models"
	"github.com/gin-gonic/gin"
)

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage:": "Could not get event by id"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorised to delete this event"})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage:": "Could not delete event by id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messsage:": "Event deleted successfully"})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage:": "Could not get event by id"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"messages": "Not correct user to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse user data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage:": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messsage:": "Event updated successfully"})

}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage:": "Could not fetch events"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not part the event"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId
	fmt.Printf("Event userid set to %v", userId)
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage:": "Could not Save() events"})

	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}
