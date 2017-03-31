package myflyingbox

import "time"

// Shipper is a shipper
type Shipper struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`

	Name    string   `json:"name,omitempty"`
	Company string   `json:"company,omitempty"`
	Street  []string `json:"street,omitempty"`
	State   string   `json:"state,omitempty"`
	Phone   string   `json:"phone,omitempty"`

	// CollectionDate is used to request collection date from a shipper when placing an order
	CollectionDate *time.Time `json:"collection_date,omitempty"`
}
