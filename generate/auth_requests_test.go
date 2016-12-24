package generate

import (
	"butter/database"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/assert"
)

func TestLoginRequestValidateReturnsFalseWithEmptyEmail(t *testing.T) {
	req := new(LoginRequest)
	req.Password = "dd"

	val := req.Validate(new(database.MockORM), &database.User{})
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"email": "The email field is required"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsFalseWithEmptyPassword(t *testing.T) {
	req := new(LoginRequest)
	req.Email = "test"

	val := req.Validate(new(database.MockORM), &database.User{})
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"password": "The password field is required"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsFalseWhenUserNotFound(t *testing.T) {
	req := new(LoginRequest)
	req.Email = "test"
	req.Password = "test"
	user := new(database.User)
	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"email": "Could not find a user with that email address"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsFalseWhenPasswordIncorrect(t *testing.T) {
	req := new(LoginRequest)
	req.Email = "test"
	req.Password = "test"

	user := new(database.User)
	user.Email = req.Email
	password, _ := bcrypt.GenerateFromPassword([]byte("test1"), 0)
	user.Password = string(password)

	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"password": "The password is incorrect"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsTrueWhenPasswordIncorrect(t *testing.T) {
	req := new(LoginRequest)
	req.Email = "test"
	req.Password = "test"

	user := new(database.User)
	user.Email = req.Email
	password, _ := bcrypt.GenerateFromPassword([]byte("test"), 0)
	user.Password = string(password)

	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.True(t, val)
}

func TestRegisterRequestValidateReturnsFalseWithEmptyEmail(t *testing.T) {
	req := new(RegisterRequest)
	req.Password = "dd"
	req.PasswordConfirmation = "dd"

	val := req.Validate(new(database.MockORM), &database.User{})
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"email": "The email field is required"}, req.GetErrors())
}

func TestRegisterRequestValidateReturnsFalseWithEmptyPassword(t *testing.T) {
	req := new(RegisterRequest)
	req.Email = "test"

	val := req.Validate(new(database.MockORM), &database.User{})
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"password": "The password field is required", "passwordConfirmation": "The passwordConfirmation field is required"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsFalseWhenUserExists(t *testing.T) {
	req := new(RegisterRequest)
	req.Email = "test"
	req.Password = "test"
	req.PasswordConfirmation = "test"

	user := new(database.User)
	user.Email = req.Email

	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"email": "User exists the given email address"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsFalseWhenPasswordUnderSix(t *testing.T) {
	req := new(RegisterRequest)
	req.Email = "test"
	req.Password = "test"
	req.PasswordConfirmation = "test"

	user := new(database.User)

	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"password": "The password field must over 6 charaters in length"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsFalseWhenPasswordConfirmationDoesNotMatch(t *testing.T) {
	req := new(RegisterRequest)
	req.Email = "test"
	req.Password = "testdsafsd"
	req.PasswordConfirmation = "testdddddddddd"

	user := new(database.User)
	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.False(t, val)
	assert.Equal(t, ApiErrors{"password": "The password and confirmation field are not the same"}, req.GetErrors())
}

func TestLoginRequestValidateReturnsTrueWhenFormCorrect(t *testing.T) {
	req := new(RegisterRequest)
	req.Email = "test"
	req.Password = "testdsafsd"
	req.PasswordConfirmation = "testdsafsd"

	user := new(database.User)
	mock := new(database.MockORM)

	mock.On("Where", "email = ? ", []interface{}{req.Email}).Return(mock)
	mock.On("First", user, []interface{}(nil)).Return(nil)

	val := req.Validate(mock, user)
	assert.True(t, val)
}
