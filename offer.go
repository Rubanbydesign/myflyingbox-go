package myflyingbox

// Offer is a offer struct
type Offer struct {
	ID              string           `json:"id"`
	QuoteID         string           `json:"quote_id"`
	ProductID       string           `json:"product_id"`
	Product         Product          `json:"product"`
	Price           *Price           `json:"price,omitempty"`
	PriceVAT        *Price           `json:"price_vat,omitempty"`
	TotalPrice      *Price           `json:"total_price,omitempty"`
	InsurancePrice  *Price           `json:"insurance_price,omitempty"`
	CollectionDates []CollectionDate `json:"collection_dates,omitempty"`

	// Response data.
	Insurable bool `json:"insurable"`
	Orderable bool `json:"orderable"`
}
