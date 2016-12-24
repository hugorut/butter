package generate

import (
	"butter/app"
	"butter/auth"
	"butter/database"
	"encoding/json"
	"time"

	"golang.org/x/crypto/bcrypt"

	"net/http"
)

// Login route, accepts a request which holds
//   - email
//   - password
// This handler validates the request and checks the password if checks
// pass then a jwt token is generated and passed back to the client
func Login(app *app.App) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var loginRequest LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			app.WriteValidationResponse(w, &ValidationErrorResponse{map[string]string{"json": "The request was in the incorrect format"}})
			return
		}

		var user database.User
		if loginRequest.Validate(app.ORM, &user) == false {
			app.WriteValidationResponse(w, &ValidationErrorResponse{loginRequest.GetErrors()})
			return
		}

		gen := auth.JWTGenerator{auth.GetSecret()}
		json.NewEncoder(w).Encode(TokenResponse{gen.GenerateToken(user)})
	})
}

// register handler for the register route, requests takes
//  - email
//  - password
//  - passwordConfirmation
// The request is validated and a new user is created and tokenised back to the client
func Register(app *App) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var registerRequest RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&registerRequest)
		if err != nil {
			app.WriteValidationResponse(w, &ValidationErrorResponse{map[string]string{"json": "The request was in the incorrect format"}})
			return
		}

		var user database.User
		if registerRequest.Validate(app.ORM, &user) == false {
			app.WriteValidationResponse(w, &ValidationErrorResponse{registerRequest.GetErrors()})
			return
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 0)
		user.Email = registerRequest.Email
		user.Password = string(password)
		user.CreatedAt = time.Now()

		app.ORM.Create(&user)

		gen := auth.JWTGenerator{auth.GetSecret()}
		json.NewEncoder(w).Encode(TokenResponse{gen.GenerateToken(user)})
	})
}

// Identify a user from the passed token, this route needs to be passed
// through the JWTProtected middleware which sets the user context
func (app *App) Identify(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	json.NewEncoder(w).Encode(user)
}
