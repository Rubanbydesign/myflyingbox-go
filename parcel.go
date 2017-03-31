package myflyingbox

// Parcel is a parcel
type Parcel struct {
	Weight float64 `json:"weight"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`

	Value              string `json:"value,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Description        string `json:"description,omitempty"`
	ShipperReference   string `json:"shipper_reference,omitempty"`
	RecipientReference string `json:"recipient_reference,omitempty"`
	CountryOfOrigin    string `json:"country_of_origin,omitempty"`
}
