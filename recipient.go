package myflyingbox

// Recipient is a recipient
type Recipient struct {
	Name       string `json:"name,omitempty"`
	Company    string `json:"company,omitempty"`
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	IsACompany bool   `json:"is_a_company,string"`

	LocationCode string `json:"location_code,omitempty"`
}
