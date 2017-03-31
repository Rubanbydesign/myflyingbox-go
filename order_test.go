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
	o, err := api.CancelOrder(ctx, &Order{ID: "invalid"})

	assert.Nil(t, o)
	assert.Error(t, err)
	assert.EqualError(t, err, "id is invalid")
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
