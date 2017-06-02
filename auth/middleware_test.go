package auth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestMiddleWareChainIsCalledCorrectly(t *testing.T) {
	i := 1

	middleWareOne := func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if i != 1 {
				t.Errorf("Expected middleware func to be called at index [%i] instead was called at [%i]", 1, i)
			}

			i += 1

			next(w, r)
		})
	}

	middleWareTwo := func(next http.HandlerFunc) http.HandlerFunc {
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

	middled := Middled(handler, middleWareOne, middleWareTwo)

	rr := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	middled(rr, r)
}

func TestJWTauthenticationMiddlewareSetsUser(t *testing.T) {
	testSig := "test"

	sub := 1

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
	})

	os.Setenv("JWT_SECRET", testSig)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(testSig))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/businesses", nil)

	r.Header.Set("Authorization", "Bearer "+tokenString)

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("sub")

		if id != float64(sub) {
			t.Errorf("Returned user is not equal to one expected: %s \n got: %s", id, sub)
		}
	}

	middled := Middled(handler, JWTProtected)
	middled(rr, r)
}
