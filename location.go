package myflyingbox

import "strings"

// Location is a location
type Location struct {
	Name   string   `json:"name"`
	Street []string `json:"street"`
	City   string   `json:"city"`

	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}

func (l Location) String() string {
	var pcs []string

	if l.Name != "" {
		pcs = append(pcs, l.Name)
	}
	for i := range l.Street {
		pcs = append(pcs, l.Street[i])
	}
	if l.City != "" {
		pcs = append(pcs, l.City)
	}
	if l.PostalCode != "" {
		pcs = append(pcs, l.PostalCode)
	}
	if l.Country != "" {
		pcs = append(pcs, l.Country)
	}

	return strings.Join(pcs, ", ")
}
