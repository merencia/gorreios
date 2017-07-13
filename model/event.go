package model

import (
	"time"
)

type Event struct {
	Date     time.Time
	Location string
	Event    string
}
