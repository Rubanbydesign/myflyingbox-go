package myflyingbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackErr(t *testing.T) {
	api := getAPI()
	_, err := api.Track(ctx, &Order{ID: "invalid"})
	assert.EqualError(t, err, "id is invalid")

	_, err = api.Track(ctx, &Quote{})
	assert.Equal(t, ErrInvalidArgumentType, err)
}
