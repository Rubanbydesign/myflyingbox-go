package myflyingbox

// Price is a price struct
type Price struct {
	Formatted     string  `json:"formatted"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount,string"`
	AmountInCents int     `json:"amount_in_cents"`
}
