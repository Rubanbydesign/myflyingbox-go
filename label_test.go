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
	req, _ := http.NewRequest("GET", uri, nil)
	req.SetBasicAuth(os.Getenv("MYFLYINGBOX_USERNAME"), os.Getenv("MYFLYINGBOX_PASSWORD"))
	client := http.DefaultClient
	resp, _ := client.Do(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
