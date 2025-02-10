package services

import (
	"errors"
	"fmt"
	"event-ticket-reservation/models"
	"event-ticket-reservation/config"
)

func GenerateSeatNumber() string {
	var lastTicket models.Ticket
	config.DB.Order("id desc").First(&lastTicket)
	return fmt.Sprintf("S-%d", lastTicket.ID+1)
}

func ReserveTicket(ticket models.Ticket) (*models.Ticket, error) {
	// Assign seat number if not provided
	if ticket.SeatNo == "" {
		ticket.SeatNo = GenerateSeatNumber()
	}

	result := config.DB.Create(&ticket)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ticket, nil
}

func GetTicket(email string) (*models.Ticket, error) {
	var ticket models.Ticket
	result := config.DB.Where("email = ?", email).First(&ticket)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ticket, nil
}

func GetAttendees(date string) ([]models.Ticket, error) {
	var attendees []models.Ticket
	result := config.DB.Where("event_date = ?", date).Find(&attendees)
	if result.Error != nil {
		return nil, result.Error
	}
	return attendees, nil
}

func CancelReservation(email, date string) error {
	result := config.DB.Where("email = ? AND event_date = ?", email, date).Delete(&models.Ticket{})
	if result.Error != nil {
		return errors.New("failed to cancel reservation")
	}
	return nil
}

func ModifyReservation(email, date, newSeat string) error {
	result := config.DB.Model(&models.Ticket{}).Where("email = ? AND event_date = ?", email, date).Update("seat", newSeat)
	if result.Error != nil {
		return errors.New("failed to modify seat reservation")
	}
	return nil
}