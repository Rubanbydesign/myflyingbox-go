package myflyingbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareRequest(t *testing.T) {

	req, err := prepareRequest(Quote{ID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", req.Quote.ID)

	req, err = prepareRequest(&Quote{ID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", req.Quote.ID)

	req, err = prepareRequest(Order{OfferID: "foobar"})
	assert.Equal(t, "foobar", req.Order.OfferID)
	assert.NoError(t, err)

	req, err = prepareRequest(&Order{OfferID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", req.Order.OfferID)

	req, err = prepareRequest("a string")
	assert.Nil(t, req)
	assert.Error(t, err)

	api := getAPI()
	assert.Error(t, api.Post(ctx, "/", "", nil))
}
