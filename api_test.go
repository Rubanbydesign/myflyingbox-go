package myflyingbox

import (
	"context"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
	o   *Order
)

func getAPI() *API {
	return New(os.Getenv("MYFLYINGBOX_USERNAME"), os.Getenv("MYFLYINGBOX_PASSWORD")).SetTestMode(true)
}

func TestSetTestMode(t *testing.T) {
	api := getAPI()

	api.SetTestMode(false)
	assert.Equal(t, ProductionURL, api.baseURL)

	api.SetTestMode(true)
	assert.Equal(t, TestURL, api.baseURL)
}
