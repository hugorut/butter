package application

import (
	"butter/database"

	"golang.org/x/crypto/bcrypt"
)

// Login requests which takes an email and password
type LoginRequest struct {
	BaseRequest
	Email    string `json:"email"`
	Password string `json:"password"`
}

// validate the login request
//  - email should be present
//  - password should be correct to unhashed version
func (l *LoginRequest) Validate(db database.ORM, user *database.User) bool {
	l.Errors = make(ApiErrors)

	m := l.NotEmpty("email", l.Email)
	p := l.NotEmpty("password", l.Password)

	if !m || !p {
		return false
	}

	db.Where("email = ? ", l.Email).First(user)

	// if we can't find the user return false
	if user.Email != l.Email {
		l.Errors["email"] = "Could not find a user with that email address"
		return false
	}

	// if the password is not the correct has invalidate the request
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password)); err != nil {
		l.Errors["password"] = "The password is incorrect"
		return false
	}

	return true
}

// register request composed of base request in order to hold validate
// functionality
type RegisterRequest struct {
	BaseRequest
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

// validate the register request
//  - email should be present and not exist
//  - password should be confirmed
//  sets errors property on the base request if failures
func (r *RegisterRequest) Validate(db database.ORM, user *database.User) bool {
	r.Errors = make(ApiErrors)

	m := r.NotEmpty("email", r.Email)
	p := r.NotEmpty("password", r.Password)
	c := r.NotEmpty("passwordConfirmation", r.PasswordConfirmation)

	if !m || !p || !c {
		return false
	}

	db.Where("email = ? ", r.Email).First(user)

	// if we can't find the user return false
	if user.Email == r.Email {
		r.Errors["email"] = "User exists the given email address"
		return false
	}

	// check that the password length is adequate
	if len(r.Password) < 6 {
		r.Errors["password"] = "The password field must over 6 charaters in length"
		return false
	}

	// check the password and confirmation are the same
	if r.Password != r.PasswordConfirmation {
		r.Errors["password"] = "The password and confirmation field are not the same"
		return false
	}

	return true
}
