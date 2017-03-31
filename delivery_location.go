package myflyingbox

// DeliveryLocation is a MyFlyingBox delivery location
type DeliveryLocation struct {
	Company      string   `json:"company"`
	Street       []string `json:"street"`
	PostalCode   string   `json:"postal_code"`
	City         string   `json:"city"`
	OpeningHours []OpeningTime
}
