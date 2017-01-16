package auth

import (
	"github.com/hugorut/butter/data"
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestGenerateReturnsClaimsWithUserId(t *testing.T) {
	var user database.User
	user.ID = 1

	prior_token := os.Getenv("JWT_SECRET")
	defer func() {
		os.Setenv("JWT_SECRET", prior_token)
	}()
	secret := "secdfldjslafjdsaf8s7fd7asd89f9a7fdsret"
	os.Setenv("JWT_SECRET", secret)

	gen := JWTGenerator{
		GetSecret(),
	}

	str := gen.GenerateToken(user)
	token, _ := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"] != float64(user.ID) {
			t.Errorf("User id not set correctly for JWT \n expected: 1 \n got: %s", claims["sub"])
		}
	} else {
		t.Error("Token is not valid")
	}
}
