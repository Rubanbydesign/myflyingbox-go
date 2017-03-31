package myflyingbox

// OpeningTime provides opening time details for a location
type OpeningTime struct {
	Day   int    `json:"day"` // 1-7 Mon-Sun
	Hours string `json:"hours"`
}
