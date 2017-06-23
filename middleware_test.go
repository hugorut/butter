package butter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleWareChainIsCalledCorrectly(t *testing.T) {
	i := 1

	middleWareOne := func(next http.HandlerFunc, app *App) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if i != 1 {
				t.Errorf("Expected middleware func to be called at index [%i] instead was called at [%i]", 1, i)
			}

			i += 1

			next(w, r)
		})
	}

	middleWareTwo := func(next http.HandlerFunc, app *App) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if i != 2 {
				t.Errorf("Expected middleware func to be called at index [%i] instead was called at [%i]", 2, i)
			}

			i += 1

			next(w, r)
		})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if i != 3 {
			t.Errorf("Expected middleware func to be called at index [%i] instead was called at [%i]", 3, i)
		}
	}

	middled := Middled(handler, &App{}, middleWareOne, middleWareTwo)

	rr := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	middled(rr, r)
}
