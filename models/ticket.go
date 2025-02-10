package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	EventDate string `json:"event_date"`
	SeatNo    string `json:"seat_no"`
}
