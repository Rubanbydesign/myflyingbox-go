// +build appengine

package myflyingbox

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// getClient (the GAE version) should return a urlfetch.Client
func (a *API) getClient(ctx context.Context) *http.Client {
	if a.Client != nil {
		return a.Client
	}
	return urlfetch.Client(ctx)
}

// New creates a new API client. If in Google App Engine development
// envirnonment, it will automatically be set to test mode
func New(login, pwd string) *API {
	a := &API{username: login, password: pwd}
	a.SetTestMode(appengine.IsDevAppServer())
	return a
}
