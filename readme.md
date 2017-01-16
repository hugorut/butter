#Butter
An api micro framework for golang, which focuses on ease of use and extensibility.

![](http://i.imgur.com/TGmKgMi.png?2)

##Work In Progress
This project is currently in progress and the SDK will change dramatically. Use with caution.

##Getting Started
You can get up and running with butter in no time at all.

* install butter using go get
`go get github.com/hugorut/butter`

* add a routing file in your package which defines a `routes` variable

```go
package main

import (
	"github.com/hugorut/butter"
	"github.com/hugorut/butter/auth"
)

var routes []butter.ApplicationRoute = []butter.ApplicationRoute{
	{
		Name:       "Index",
		Method:     "GET",
		URI:        "/",
		Func:       Index,
		Middleware: auth.Chain{},
	},
}
```

* define a butter handler, in this case our `Index` function, the signature is as follows:

```go

func Index(app *butter.App) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello World!")

		w.WriteHeader(200)
	})
}
```

* in your main function run the butter server with the specified routes

```go
package main

import (
	"butter"
)

func main() {
	butter.Serve(routes)
}
```

* give yourself a pat on the back, you just implemented your first page for your API.

##Configuration
Butter uses environment variables to give you the power to customise your butter application.
Butter uses `godotenv` to load environment variables from your `.env` file at the root of your project so that.


