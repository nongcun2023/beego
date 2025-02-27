// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package auth provides handlers to enable basic auth support.
// Simple Usage:
//
//	import(
//		"github.com/beego/beego/v2"
//		"github.com/beego/beego/v2/server/web/filter/auth"
//	)
//
//	func main(){
//		// authenticate every request
//		beego.InsertFilter("*", beego.BeforeRouter,auth.Basic("username","secretpassword"))
//		beego.Run()
//	}
//
// Advanced Usage:
//
//	func SecretAuth(username, password string) bool {
//		return username == "astaxie" && password == "helloBeego"
//	}
//	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
//	beego.InsertFilter("*", beego.BeforeRouter,authPlugin)
package auth

import (
	"net/http"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/auth"
)

// Basic is the http basic auth
func Basic(username string, password string) beego.FilterFunc {
	return func(c *context.Context) {
		f := auth.Basic(username, password)
		f((*beecontext.Context)(c))
	}
}

// NewBasicAuthenticator return the BasicAuth
func NewBasicAuthenticator(secrets SecretProvider, realm string) beego.FilterFunc {
	f := auth.NewBasicAuthenticator(auth.SecretProvider(secrets), realm)
	return func(c *context.Context) {
		f((*beecontext.Context)(c))
	}
}

// SecretProvider is the SecretProvider function
type SecretProvider auth.SecretProvider

// BasicAuth store the SecretProvider and Realm
type BasicAuth auth.BasicAuth

// CheckAuth Checks the username/password combination from the request. Returns
// either an empty string (authentication failed) or the name of the
// authenticated user.
// Supports MD5 and SHA1 password entries
func (a *BasicAuth) CheckAuth(r *http.Request) string {
	return (*auth.BasicAuth)(a).CheckAuth(r)
}

// RequireAuth http.Handler for BasicAuth which initiates the authentication process
// (or requires reauthentication).
func (a *BasicAuth) RequireAuth(w http.ResponseWriter, r *http.Request) {
	(*auth.BasicAuth)(a).RequireAuth(w, r)
}
