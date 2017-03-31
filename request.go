package myflyingbox

// Request is a struct used for sending requests of all types since
// request elements need to be wrapped in a root element of the same
// name
type Request struct {
	Quote *Quote `json:"quote,omitempty"`
	Order *Order `json:"order,omitempty"`
}

func prepareRequest(el ...interface{}) (*Request, error) {
	var req Request
	for i := range el {
		switch el[i].(type) {
		case Quote:
			q := el[i].(Quote)
			req.Quote = &q
		case *Quote:
			req.Quote = el[i].(*Quote)
		case Order:
			o := el[i].(Order)
			req.Order = &o
		case *Order:
			if o := el[i].(*Order); o != nil {
				req.Order = el[i].(*Order)
			}
		default:
			return nil, ErrInvalidArgumentType
		}
	}
	return &req, nil
}
