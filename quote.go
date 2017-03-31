package myflyingbox

// Quote is a MyFlyingBox quote
type Quote struct {
	Shipper   Shipper   `json:"shipper"`
	Recipient Recipient `json:"recipient"`
	Parcels   []Parcel  `json:"parcels"`
	Offers    []Offer   `json:"offers"`

	ID string `json:"id,omitempty"`

	Ordered bool `json:"ordered"`
}
