package myflyingbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceOrderErr(t *testing.T) {
	api := getAPI()
	testOrder := Order{}
	res, err := api.PlaceOrder(ctx, &testOrder)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestOrderErr(t *testing.T) {
	api := getAPI()
	res, err := api.GetOrder(ctx, &Order{ID: "invalid"})
	assert.Nil(t, res)
	assert.EqualError(t, err, "id is invalid")
}

func TestCancelOrderErr(t *testing.T) {
	api := getAPI()
	_, err := api.CancelOrder(ctx, &Order{ID: "invalid"})
	assert.EqualError(t, err, "id is invalid")

	_, err = api.CancelOrder(ctx, &Quote{})
	assert.Equal(t, ErrInvalidArgumentType, err)

	_, err = api.CancelOrder(ctx, &Order{ID: ""})
	assert.Equal(t, ErrNoOrderID, err)
}

func TestGetOrderID(t *testing.T) {
	orderID, err := getOrderID("foobar")
	assert.NoError(t, err)
	assert.Equal(t, "foobar", orderID)

	orderID, err = getOrderID(&orderID)
	assert.NoError(t, err)
	assert.Equal(t, "foobar", orderID)

	orderID, err = getOrderID(Order{ID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", orderID)

	orderID, err = getOrderID(&Order{ID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", orderID)

	_, err = getOrderID(&Quote{})
	assert.Equal(t, ErrInvalidArgumentType, err)
}

func TestGetOrderErr(t *testing.T) {
	api := getAPI()
	_, err := api.GetOrder(ctx, "invalid")
	assert.Error(t, err)

	_, err = api.GetOrder(ctx, &Quote{})
	assert.Error(t, err)
}

func TestOrder(t *testing.T) {
	// Quote
	api := getAPI()
	q := Quote{
		Shipper:   Shipper{Country: "US", PostalCode: "11201", City: "Brooklyn"},
		Recipient: Recipient{IsACompany: false, Country: "US", PostalCode: "11201", City: "Brooklyn"},
		Parcels: []Parcel{
			Parcel{Weight: 1.0, Length: 10, Width: 10, Height: 10},
		},
	}

	quote, err := api.RequestQuote(ctx, &q)
	if !assert.NoError(t, err) {
		return
	}

	// Place
	o := Order{
		OfferID:        quote.Offers[0].ID,
		InsureShipment: false,
		Shipper: Shipper{
			Name:   "test shipment",
			Street: "1 main st.",
			City:   "brooklyn",
			State:  "NY",
			Phone:  "+33123456789",
			Email:  "test@example.org",
		},
		Recipient: Recipient{
			Name:   "test recipient",
			Street: "1 main st.",
			City:   "brooklyn",
			State:  "ny",
			Email:  "test@example.org",
			Phone:  "+33123456789",
		},
		Parcels: []Parcel{
			Parcel{
				CustomerReference:  "test cust. ref",
				ShipperReference:   "test ship. ref.",
				RecipientReference: "test rcpt. ref.",
				Value:              10,
				Currency:           "USD",
				Description:        "test parcel desc.",
				CountryOfOrigin:    "US",
				InsuredValue:       10,
				InsuredCurrency:    "USD",
			},
		},
	}

	order, err := api.PlaceOrder(ctx, &o)
	if !assert.NoError(t, err) {
		return
	}

	// Get order details.
	_, err = api.GetOrder(ctx, order)
	if !assert.NoError(t, err) {
		return
	}

	// Track
	track, err := api.Track(ctx, order)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotNil(t, track)
	assert.Len(t, track, 0)

	// Cancel
	_, err = api.CancelOrder(ctx, order)
	assert.NoError(t, err)
}
