package myflyingbox

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiStdGetClient(t *testing.T) {
	api := getAPI()
	c := &http.Client{}
	api.Client = c
	assert.Equal(t, c, api.getClient(context.Background()))
}
