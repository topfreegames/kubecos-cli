// mystack api
// https://github.com/topfreegames/mystack
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package api

import (
	"github.com/topfreegames/mystack/mystack-cli/errors"
	"github.com/topfreegames/mystack/mystack-cli/models"
	"net/http"
)

//OAuthCallbackHandler handles the callback after user approves/deny auth
type OAuthCallbackHandler struct {
	App *App
}

//ServeHTTP method
func (o *OAuthCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")
	l := loggerFromContext(r.Context())

	err := models.SaveAccessToken(state, code, o.App.Login.OAuthState)
	if err != nil {
		if err, ok := err.(*errors.OAuthError); ok {
			l.Error(err.Serialize())
		}

		l.Error(err)
	}

	o.App.ServerControl.CloseServer <- true
}