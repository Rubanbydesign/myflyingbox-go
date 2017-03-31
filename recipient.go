package myflyingbox

// Recipient is a recipient
type Recipient struct {
	Name       string   `json:"name"`
	Company    string   `json:"company"`
	Street     []string `json:"street"`
	City       string   `json:"city"`
	State      string   `json:"state"`
	PostalCode string   `json:"postal_code"`
	Country    string   `json:"country"`
	Phone      string   `json:"phone"`
	Email      string   `json:"email"`
	IsACompany bool     `json:"is_a_company,string"`

	LocationCode string `json:"location_code,omitempty"`
}
