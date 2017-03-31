package myflyingbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackErr(t *testing.T) {
	api := getAPI()
	trk, err := api.Track(ctx, &Order{ID: "invalid"})
	assert.Nil(t, trk)
	assert.EqualError(t, err, "id is invalid")
}
