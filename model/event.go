package model

import (
	"time"
)

type Event struct {
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
	Event    string    `json:"event"`
}
