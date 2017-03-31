// +build !appengine

package myflyingbox

import (
	"net"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

var (
	defaultClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
)

// getClient returns an HTTP client. If not has been set up with the client,
// it sets a default with 5 second timeout.
func (a *API) getClient(ctx context.Context) *http.Client {
	if a.Client != nil {
		return a.Client
	}
	return defaultClient
}

// New creates a new API client. If in Google App Engine development
// envirnonment, it will automatically be set to test mode
func New(login, pwd string) *API {
	a := &API{username: login, password: pwd}
	a.SetTestMode(false)
	return a
}
