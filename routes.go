package main

import (
	"butter/application"
	"butter/auth"
	"butter/routing"
)

var routes []routing.ApplicationRoute = []routing.ApplicationRoute{
	routing.ApplicationRoute{
		Name:       "Login",
		Method:     "POST",
		URI:        "/login",
		Func:       application.Login,
		Middleware: auth.Chain{auth.CrossOrigin},
	},
	routing.ApplicationRoute{
		Name:       "Register",
		Method:     "POST",
		URI:        "/register",
		Func:       application.Register,
		Middleware: auth.Chain{auth.CrossOrigin},
	},
}
