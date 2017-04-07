// Package myflyingbox implements a Go REST client for the MyFlyingBox API
package myflyingbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/context"
)

const (
	// ProductionURL is the production API base endpoint
	ProductionURL = "https://api.myflyingbox.com/v2"
	// TestURL is the production API base endpoint
	TestURL = "https://test.myflyingbox.com/v2"
)

// Error messages
var (
	// ErrInvalidArgumentType is returns when a argument which has an interface
	// type does not recieve a value it can use
	ErrInvalidArgumentType = errors.New("invalid argument type")

	ErrNoOrderID = errors.New("no order id")
	ErrNoQuoteID = errors.New("no quote id")

	captureRespBody bool
)

// API is a myflyingbox.com API client for Go
type API struct {
	baseURL            string
	username, password string

	// Client is the http client used for requests. If nil, an appropriate default is used
	Client *http.Client
}

// Do sends the request and marshals the response into result
func (a *API) Do(ctx context.Context, req *http.Request, result interface{}) error {
	req.SetBasicAuth(a.username, a.password)
	resp, err := a.getClient(ctx).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// For testing only, allows inspection of response body by reading and reset.
	var bodyBytes []byte
	if captureRespBody {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	var apiResp Response
	if err = json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}
	if err = apiResp.Error(); err != nil {
		log.Printf("Body: %s", string(bodyBytes))
		return err
	}
	// Could be done better here, but marshal the apiResp into json, then decode
	// back into result.
	if result != nil {
		data, err := json.Marshal(apiResp.Data)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(data, result); err != nil {
			return err
		}
	}

	return nil
}

// Post sends a POST request with the JSON encoded data to the urlStr and marshals response into result
func (a *API) Post(ctx context.Context, urlStr string, data, result interface{}) error {

	// Requests need to be enclosed in a "root" element per the API
	// docs, so we do that here.
	payload, err := prepareRequest(data)
	if err != nil {
		return err
	}

	// Marshal the data. Since the payload was already formatted, there won't be any errors encoding.
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	// Create the request
	req, _ := http.NewRequest("POST", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), buf)
	req.Header.Set("Content-Type", "application/json")
	return a.Do(ctx, req, result)
}

// Put is for delete requests, so only the urlStr is implemented
func (a *API) Put(ctx context.Context, urlStr string, result interface{}) error {
	req, _ := http.NewRequest("PUT", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), nil)
	return a.Do(ctx, req, result)
}

// Get sends a GET request to the urlStr and marshals response into result
func (a *API) Get(ctx context.Context, urlStr string, result interface{}) error {
	req, _ := http.NewRequest("GET", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), nil)
	return a.Do(ctx, req, result)
}

// SetURL sets the API base url
func (a *API) SetURL(urlStr string) *API {
	a.baseURL = urlStr
	return a
}

// SetTestMode toggles mode to test
func (a *API) SetTestMode(test bool) *API {
	if test {
		a.SetURL(TestURL)
		return a
	}
	a.SetURL(ProductionURL)
	return a
}

// CancelOrder cancels the given order.
func (a *API) CancelOrder(ctx context.Context, o interface{}) (*Order, error) {
	var result Order
	orderID, err := getOrderID(o)
	if err != nil {
		return nil, err
	}
	if err := a.Put(ctx, fmt.Sprintf("/orders/%s/cancel", orderID), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// LabelPDF returns the label PDF as a byte slice
func (a *API) LabelPDF(ctx context.Context, o interface{}) ([]byte, error) {
	orderID, err := getOrderID(o)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/orders/%s/labels", a.baseURL, orderID), nil)
	req.SetBasicAuth(a.username, a.password)
	resp, err := a.getClient(ctx).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

// Track returns tracking information about the given order
func (a *API) Track(ctx context.Context, o interface{}) (list []Tracking, err error) {

	orderID, err := getOrderID(o)
	if err != nil {
		return nil, err
	}

	if err := a.Get(ctx, fmt.Sprintf("/orders/%s/tracking", orderID), &list); err != nil {
		return nil, err
	}
	return
}

// GetOrder returns order info for the given order
func (a *API) GetOrder(ctx context.Context, o interface{}) (*Order, error) {

	orderID, err := getOrderID(o)
	if err != nil {
		return nil, err
	}
	var result Order
	if err := a.Get(ctx, "/orders/"+orderID, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PlaceOrder creates an order and returns the results
func (a *API) PlaceOrder(ctx context.Context, o *Order) (*Order, error) {
	var result Order
	if err := a.Post(ctx, "/orders", o, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestQuote places a quote request and returns the result
func (a *API) RequestQuote(ctx context.Context, q *Quote) (*Quote, error) {
	var quote Quote
	if err := a.Post(ctx, "/quotes", q, &quote); err != nil {
		return nil, err
	}
	return &quote, nil
}

// RetrieveQuote retrieves details for an existing quote.
func (a *API) RetrieveQuote(ctx context.Context, q interface{}) (*Quote, error) {

	quoteID, err := getQuoteID(q)
	if err != nil {
		return nil, err
	}

	var quote Quote
	if err := a.Get(ctx, "/quotes/"+quoteID, &quote); err != nil {
		return nil, err
	}
	return &quote, nil
}

func getQuoteID(q interface{}) (quoteID string, err error) {
	switch q.(type) {
	case string:
		quoteID = q.(string)
	case *string:
		quoteID = *q.(*string)
	case Quote:
		quoteID = q.(Quote).ID
	case *Quote:
		quoteID = q.(*Quote).ID
	default:
		err = ErrInvalidArgumentType
	}
	if err == nil && quoteID == "" {
		err = ErrNoQuoteID
	}
	return
}

func getOrderID(o interface{}) (orderID string, err error) {
	switch o.(type) {
	case string:
		orderID = o.(string)
	case *string:
		orderID = *o.(*string)
	case Order:
		orderID = o.(Order).ID
	case *Order:
		orderID = o.(*Order).ID
	default:
		err = ErrInvalidArgumentType
	}
	if err == nil && orderID == "" {
		err = ErrNoOrderID
	}
	return
}
