package myflyingbox

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelURLErr(t *testing.T) {
	api := getAPI()
	uri := api.LabelURL(ctx, &Order{ID: "foo"})
	req, err := http.NewRequest("GET", uri, nil)
	if !assert.NoError(t, err) {
		return
	}
	req.SetBasicAuth(os.Getenv("MYFLYINGBOX_USERNAME"), os.Getenv("MYFLYINGBOX_PASSWORD"))
	client := http.DefaultClient
	resp, err := client.Do(req)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
