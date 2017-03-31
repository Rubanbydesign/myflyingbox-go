package myflyingbox

// Tracking is a tracking response
type Tracking struct {
	ParcelIndex int     `json:"parcel_index"`
	Events      []Event `json:"events"`
}
