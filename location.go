package myflyingbox

// Location is a location
type Location struct {
	Street []string `json:"street"`
	City   string   `json:"city"`

	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}
