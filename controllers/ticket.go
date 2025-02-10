package controllers

import (
	"event-ticket-reservation/models"
	"event-ticket-reservation/services"
	"event-ticket-reservation/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReserveTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}
	if ticket.SeatNo == "" {
		ticket.SeatNo = services.GenerateSeatNumber()
	}

	reservedTicket, err := services.ReserveTicket(ticket)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to reserve ticket: "+err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, gin.H{
		"success": true,
		"message": "Ticket reserved successfully",
		"data":    reservedTicket,
	})
}


func GetTicket(c *gin.Context) {
	email := c.Query("email")
	
	ticket, err := services.GetTicket(email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.JSONResponse(c, http.StatusOK, ticket)
}

func GetAttendees(c *gin.Context) {
	date := c.Query("date")
	attendees, err := services.GetAttendees(date)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(c, http.StatusOK, attendees)
}

func CancelReservation(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Date  string `json:"event_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := services.CancelReservation(req.Email, req.Date); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(c, http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}

func ModifyReservation(c *gin.Context) {
	var req struct {
		Email   string `json:"email"`
		Date    string `json:"event_date"`
		NewSeat string `json:"new_seat"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := services.ModifyReservation(req.Email, req.Date, req.NewSeat); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(c, http.StatusOK, gin.H{"message": "Seat reservation modified successfully"})
}
