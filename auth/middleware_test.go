package auth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hugorut/butter"
)

func TestGenerateReturnsClaimsWithUserId(t *testing.T) {
	id := 1

	prior_token := os.Getenv("JWT_SECRET")
	defer func() {
		os.Setenv("JWT_SECRET", prior_token)
	}()
	secret := "secdfldjslafjdsaf8s7fd7asd89f9a7fdsret"
	os.Setenv("JWT_SECRET", secret)

	gen := JWTGenerator{
		GetSecret(),
	}

	str := gen.GenerateToken(id)
	token, _ := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"] != float64(id) {
			t.Errorf("User id not set correctly for JWT \n expected: 1 \n got: %s", claims["sub"])
		}
	} else {
		t.Error("Token is not valid")
	}
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

	middled := butter.Middled(handler, &butter.App{}, JWTProtected)
	middled(rr, r)
}
