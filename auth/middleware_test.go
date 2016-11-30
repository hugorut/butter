package auth

import (
	"butter/database"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestMiddleWareChainIsCalledCorrectly(t *testing.T) {
	i := 1
	db := new(database.MockORM)

	middleWareOne := func(db database.ORM, next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if i != 1 {
				t.Errorf("Expected middleware func to be called at index [%i] instead was called at [%i]", 1, i)
			}

			i += 1

			next(w, r)
		})
	}

	middleWareTwo := func(db database.ORM, next http.HandlerFunc) http.HandlerFunc {
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

	middled := Middled(db, handler, middleWareOne, middleWareTwo)

	rr := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	middled(rr, r)
}

func TestJWTauthenticationMiddlewareSetsUser(t *testing.T) {
	testSig := "test"

	mockUser := database.User{}
	mockUser.ID = 1
	mockUser.Email = "hugorut@gmail.com"

	db := &database.MockORMSetsUser{
		new(database.MockORM),
		mockUser.Email,
		"",
		uint64(mockUser.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": mockUser.ID,
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
		returnedUser := r.Context().Value("user")
		ps := reflect.ValueOf(returnedUser)
		id := ps.FieldByName("ID").Interface()

		if id != mockUser.ID {
			t.Errorf("Returned user is not equal to one expected: %s \n got: %s", id, mockUser.ID)
		}
	}

	middled := Middled(db, handler, JWTProtected)
	middled(rr, r)
}
