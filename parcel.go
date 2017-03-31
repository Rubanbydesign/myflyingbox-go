package myflyingbox

// Parcel is a parcel
type Parcel struct {
	Weight float64 `json:"weight,omitempty"`
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`

	Value              int64  `json:"value,string,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Description        string `json:"description,omitempty"`
	ShipperReference   string `json:"shipper_reference,omitempty"`
	RecipientReference string `json:"recipient_reference,omitempty"`
	CustomerReference  string `json:"customer_reference,omitempty"`
	CountryOfOrigin    string `json:"country_of_origin,omitempty"`

	InsuredValue    float64 `json:"insured_value,omitempty"`
	InsuredCurrency string  `json:"insured_currency,omitempty"`
}
