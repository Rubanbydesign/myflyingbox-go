package myflyingbox

// Request is a struct used for sending requests of all types since
// request elements need to be wrapped in a root element of the same
// name
type Request struct {
	Quote *Quote `json:"quote,omitempty"`
}

func prepareRequest(el ...interface{}) *Request {
	var req Request
	for i := range el {
		switch el[i].(type) {
		case Quote:
			q := el[i].(Quote)
			req.Quote = &q
		case *Quote:
			req.Quote = el[i].(*Quote)
		}
	}
	return &req
}
