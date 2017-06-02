package butter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hugorut/butter/sys"
	"github.com/hugorut/butter/auth"
)

func TestGorillaRouter_AddRoutes_RegistersRoutesForRequest(t *testing.T) {
	var called bool
	logger := &sys.TestLogger{}
	r := NewGorillaRouter(logger)

	routes := Routes{
		{
			Method: "GET",
			URI:    "/test/url",
			HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
			}),
		},
	}

	r.AddRoutes(routes)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://test.com/test/url", nil)

	r.ServeHTTP(res, req)

	if !called {
		t.Error("Failed asserting that added route was called successfully")
	}
}

func TestGorillaRouter_ServeHTTP_CatchesPanic_IfFailGracefulEnabled(t *testing.T) {
	logger := &sys.TestLogger{}
	r := NewGorillaRouter(logger)

	routes := Routes{
		{
			Method: "GET",
			URI:    "/test/url",
			HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("test panic")
			}),
		},
	}

	r.AddRoutes(routes)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://test.com/test/url", nil)

	r.ServeHTTP(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("Failed asserting that the response was a internal server error, instead was %v", res.Code)
	}

	if !logger.AssertCalled(sys.CRITICAL, "Panic recovered from handler\nmethod: GET\nreq: /test/url\nname: github.com/hugorut/butter.TestGorillaRouter_ServeHTTP_CatchesPanic_IfFailGracefulEnabled.func1\nline: 46\nerr: test panic") {
		t.Errorf("Logger was not called with panic recovery message")
	}

}

func TestApplyRoutes_TranslatesApplicationRoutes_ToRoutes(t *testing.T) {
	var calledApplicationRoute bool
	var calledBaseRoute bool

	app, _ := NewMockApplication()

	routes := []ApplicationRoute{
		{
			Name: "test",
			Method: "GET",
			URI: "/test/uri",
			Func: func (*App) http.HandlerFunc {
				calledApplicationRoute = true
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
					calledBaseRoute = true
				})
			},
		},
	}

	actual := ApplyRoutes(app, routes, auth.SkipMiddleware)

	route := actual[0]
	if route.URI != "/test/uri" {
		t.Errorf("Failed asserting route URI /test/uri, instead was %s", route.URI)
	}

	if route.Method != "GET" {
		t.Errorf("Failed asserting route method GET, instead was %s", route.Method)
	}

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://test.com/test/uri", nil)

	if !calledApplicationRoute {
		t.Error("Failed asserting that application wrapping function called")
	}

	route.HandlerFunc(res, req)

	if ! calledBaseRoute {
		t.Error("Failed asserting that the correct base handler func called")
	}

}
