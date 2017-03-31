// Package myflyingbox implements a Go REST client for the MyFlyingBox API
package myflyingbox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// ProductionURL is the production API base endpoint
	ProductionURL = "https://api.myflyingbox.com/v2"
	// TestURL is the production API base endpoint
	TestURL = "https://test.myflyingbox.com/v2"
)

var (
	// ErrInvalidArgumentType is returns when a argument which has an interface
	// type does not recieve a value it can use
	ErrInvalidArgumentType = errors.New("invalid argument type")
)

// Send interface is used to define custom data interfaces for sending data to the API.
// Some (if not all) structs need to be wrapped in a "root" element for POSTing, so structs
// which require that can implement the Sender() interface to return a custom version of
// themselves which satisfies the API
type Send interface {
	Send() interface{}
}

// Form interface is used to define custom data interfaces for sending data to the API.
// Some (if not all) structs need to be wrapped in a "root" element for POSTing, so structs
// which require that can implement the Form() interface to return a custom set of key/Values
// pairs for posting to the API as a encoded form.
type Form interface {
	Form() url.Values
}

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

	// TODO For testing only, remove when done
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Reset the body again.
	fmt.Printf("Body: %s", string(bodyBytes))
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var apiResp Response
	if err = json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}
	if err = apiResp.Error(); err != nil {
		return err
	}
	// Could be done better here, but marshal the apiResp into json, then decode
	// back into result.
	data, err := json.Marshal(apiResp.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// PostJSON sends a POST request with the JSON encoded data to the urlStr and marshals response into result
func (a *API) PostJSON(ctx context.Context, urlStr string, data, result interface{}) error {

	// If implementing the Send interface, then use that data instead.
	if send, ok := data.(Send); ok {
		data = send.Send()
	}

	// Requests need to be enclosed in a "root" element per the API
	// docs, so we do that here.
	payload := prepareRequest(data)

	// Marshal the data
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return a.Do(ctx, req, result)
}

// Post sends a POST request to the urlStr and marshals response into result
func (a *API) Post(ctx context.Context, urlStr string, data, result interface{}) error {
	var form url.Values

	if f, ok := data.(url.Values); ok {
		form = f
	}
	if fdata, ok := data.(Form); ok {
		form = fdata.Form()
	}
	if form == nil || len(form) < 1 {
		return errors.New("form is empty")
	}

	req, err := http.NewRequest("POST", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), bytes.NewBufferString(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return a.Do(ctx, req, result)
}

// Put is for delete requests, so only the urlStr is implemented
func (a *API) Put(ctx context.Context, urlStr string, result interface{}) error {
	req, err := http.NewRequest("PUT", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), nil)
	if err != nil {
		return nil
	}
	return a.Do(ctx, req, result)
}

// Get sends a GET request to the urlStr and marshals response into result
func (a *API) Get(ctx context.Context, urlStr string, result interface{}) error {
	req, err := http.NewRequest("GET", a.baseURL+"/"+strings.TrimPrefix(urlStr, "/"), nil)
	if err != nil {
		return err
	}
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

// LabelURL returns a url which can be used to download the label PDF directly.
func (a *API) LabelURL(ctx context.Context, o *Order) string {
	return fmt.Sprintf("%s/v2/orders/%s/labels", a.baseURL, o.ID)
}

// Track returns tracking information about the given order
func (a *API) Track(ctx context.Context, o interface{}) (*Tracking, error) {

	var t Tracking
	orderID, err := getOrderID(o)
	if err != nil {
		return nil, err
	}

	if err := a.Get(ctx, fmt.Sprintf("/orders/%s/tracking", orderID), &t); err != nil {
		return nil, err
	}
	return &t, nil
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
	if err := a.PostJSON(ctx, "/orders", o, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestQuote places a quote request and returns the result
func (a *API) RequestQuote(ctx context.Context, q *Quote) (*Quote, error) {
	var quote Quote
	if err := a.PostJSON(ctx, "/quotes", q, &quote); err != nil {
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
	return
}

// getRootElement gets a root element with a given type for encoding JSON requests.
func getRootElement(key string, el interface{}) map[string]interface{} {
	return map[string]interface{}{key: el}
}
