package routes

import (
	"event-ticket-reservation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/tickets/reserve", controllers.ReserveTicket)
	r.GET("/tickets", controllers.GetTicket)
	r.GET("/tickets/attendees", controllers.GetAttendees)
	r.DELETE("/tickets/cancel", controllers.CancelReservation)
}
