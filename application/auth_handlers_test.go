package application

import (
	"butter/database"
	"butter/service"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestLoginRequestRecievesErrorResponseFromEmptyBody(t *testing.T) {
	w, r := service.NewJsonPostRequest("/login", []byte(`{}`))

	mock := new(database.MockORM)
	app := &App{
		DB: mock,
	}

	app.Login(w, r)
	actual, _ := ioutil.ReadAll(w.Body)

	expected := `{"errors":{"email":"The email field is required","password":"The password field is required"}}`

	assert.Equal(t, expected, strings.TrimSpace(string(actual)))
}

func TestLoginRequestRecievesErrorResponseFromIncorrectEmail(t *testing.T) {
	email := "hugorut@gmail.com"
	password := "test"
	w, r := service.NewJsonPostRequest("/login", []byte(`{"email":"`+email+`","password":"`+password+`"}`))

	mock := new(database.MockORM)
	app := &App{
		DB: mock,
	}
	user := new(database.User)

	mock.On("Where", "email = ? ", []interface{}{email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	app.Login(w, r)
	actual, _ := ioutil.ReadAll(w.Body)

	expected := `{"errors":{"email":"Could not find a user with that email address"}}`
	assert.Equal(t, expected, strings.TrimSpace(string(actual)))
}

func TestLoginRequestRecievesErrorResponseFromIncorrectPassword(t *testing.T) {
	email := "hugorut@gmail.com"
	password := "test"
	w, r := service.NewJsonPostRequest("/login", []byte(`{"email":"`+email+`","password":"`+password+`"}`))

	mock := &database.MockORMSetsUser{
		new(database.MockORM),
		email,
		"pass",
		uint64(3),
	}

	app := &App{
		DB: mock,
	}
	user := new(database.User)

	mock.On("Where", "email = ? ", []interface{}{email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	app.Login(w, r)
	actual, _ := ioutil.ReadAll(w.Body)

	expected := `{"errors":{"password":"The password is incorrect"}}`
	assert.Equal(t, expected, strings.TrimSpace(string(actual)))
}

func TestLoginRequestRecievesSuccessFromCorrectPassword(t *testing.T) {

	prior_token := os.Getenv("JWT_SECRET")
	defer func() {
		os.Setenv("JWT_SECRET", prior_token)
	}()

	email := "hugorut@gmail.com"
	password := "test"
	w, r := service.NewJsonPostRequest("/login", []byte(`{"email":"`+email+`","password":"`+password+`"}`))

	encrypt, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	mock := &database.MockORMSetsUser{
		new(database.MockORM),
		email,
		string(encrypt),
		uint64(1),
	}

	app := &App{
		DB: mock,
	}
	user := new(database.User)

	mock.On("Where", "email = ? ", []interface{}{email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	secret := "secdfldjslafjdsaf8s7fd7asd89f9a7fdsret"
	os.Setenv("JWT_SECRET", secret)

	app.Login(w, r)

	var response map[string]string
	actual, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(actual, &response)

	if _, ok := response["token"]; !ok {
		t.Errorf("Token response could not be found got:\n %s", response)
	}

	token, _ := jwt.Parse(response["token"], func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"] != float64(1) {
			t.Errorf("User id not set correctly for JWT \n expected: 1 \n got: %s", claims["sub"])
		}
	} else {
		t.Error("Token is not valid")
	}
}

func TestRegisterRequestRecievesErrorResponseFromDifferentPasswordConfirmation(t *testing.T) {
	email := "hugorut@gmail.com"
	password := "thisisoversix"
	w, r := service.NewJsonPostRequest("/register", []byte(`{"email":"`+email+`","password":"`+password+`", "passwordConfirmation":"`+password+`dd"}`))

	mock := new(database.MockORM)

	app := &App{
		DB: mock,
	}
	user := new(database.User)

	mock.On("Where", "email = ? ", []interface{}{email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	app.Register(w, r)
	actual, _ := ioutil.ReadAll(w.Body)

	expected := `{"errors":{"password":"The password and confirmation field are not the same"}}`
	assert.Equal(t, expected, strings.TrimSpace(string(actual)))
}

func TestRegisterRequestRecievesTokenFromCorrectPassword(t *testing.T) {
	prior_token := os.Getenv("JWT_SECRET")
	defer func() {
		os.Setenv("JWT_SECRET", prior_token)
	}()

	email := "hugorut@gmail.com"
	password := "testisoversix"
	w, r := service.NewJsonPostRequest("/register", []byte(`{"email":"`+email+`","password":"`+password+`", "passwordConfirmation":"`+password+`"}`))

	encrypt, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	mock := &database.MockORMCreatesUser{
		new(database.MockORM),
		"",
		email,
		string(encrypt),
		uint64(4),
	}

	app := &App{
		DB: mock,
	}
	user := new(database.User)

	mock.On("Where", "email = ? ", []interface{}{email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	secret := "secdfldjslafjdsaf8s7fd7asd89f9a7fdsret"
	os.Setenv("JWT_SECRET", secret)

	app.Register(w, r)

	var response map[string]string
	actual, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(actual, &response)

	if _, ok := response["token"]; !ok {
		t.Errorf("Token response could not be found got:\n %s", actual)
		return
	}

	token, _ := jwt.Parse(response["token"], func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"] != float64(4) {
			t.Errorf("User id not set correctly for JWT \n expected: 1 \n got: %s", claims["sub"])
		}
	} else {
		t.Error("Token is not valid")
	}
}

func TestIdentifyRequestRecievesUserSetInContext(t *testing.T) {
	w, r := service.NewGetRequest("/identify")

	user := database.User{
		3,
		"Hugo",
		"hugorut@gmail.com",
		"password",
		time.Now(),
		[]database.Test{},
	}
	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)

	db := new(database.MockORM)
	app := &App{
		DB: db,
	}

	app.Identify(w, r)
	var returnedUser database.User

	actual, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(actual, &returnedUser)

	assert.Equal(t, user.Name, returnedUser.Name, string(actual))
	assert.Equal(t, user.Email, returnedUser.Email, string(actual))
	assert.Equal(t, "", returnedUser.Password, string(actual))
	assert.Equal(t, user.ID, returnedUser.ID, string(actual))
}
