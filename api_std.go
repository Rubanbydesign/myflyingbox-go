// +build !appengine

package myflyingbox

import (
	"net/http"

	"golang.org/x/net/context"
)

func (a *API) getClient(ctx context.Context) *http.Client {
	if a.Client != nil {
		return a.Client
	}
	return http.DefaultClient
}

// New creates a new API client. If in Google App Engine development
// envirnonment, it will automatically be set to test mode
func New(login, pwd string) *API {
	a := &API{username: login, password: pwd}
	a.SetTestMode(false)
	return a
}
