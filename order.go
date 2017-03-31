package myflyingbox

// Order is the struct used to place an order.
type Order struct {

	// Used in placing the order.

	OfferID        string  `json:"offer_id"`
	ThermalLabels  bool    `json:"thermal_labels"`
	InsureShipment bool    `json:"insure_shipment"`
	Shipper        Shipper `json:"shipper"`

	Parcels []Parcel `json:"parcels"`

	// Used when retrieving the order
	ID        string    `json:"uuid,omitempty"`
	Recipient Recipient `json:"recipient,omitempty"`
}
