package myflyingbox

// Product is a MyFlyingBox product
type Product struct {
	ID                     string `json:"id"`
	Logo                   string `json:"logo"`
	Code                   string `json:"code"`
	Name                   string `json:"name"`
	Delay                  string `json:"delay"`
	Pickup                 bool   `json:"pick_up"`
	Dropoff                bool   `json:"drop_off"`
	PresetDeliveryLocation bool   `json:"preset_delivery_location"`

	// Response vars

	ExportFrom             []string          `json:"export_from,omitempty"`
	CarrierCode            string            `json:"carrier_code,omitempty"`
	CollectionInformations map[string]string `json:"collection_informations,omitempty"`
	DeliveryInformations   map[string]string `json:"delivery_informations,omitempty"`
	Details                map[string]string `json:"details,comitempty"`
}
