package myflyingbox

import "time"

// Event is a tracking event
type Event struct {
	Code      string    `json:"code"`
	HappendAt time.Time `json:"happend_at"`
	Label     Label     `json:"label"`
	Location  Location  `json:"location"`
	Details   Details   `json:"details"`
}
