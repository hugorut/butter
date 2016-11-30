package application

import (
	"butter/database"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var UserOutOfContext = errors.New("User out of context")

func ErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func NotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func WriteValidationResponse(w http.ResponseWriter, v *ValidationErrorResponse) {
	w.WriteHeader(422)
	json.NewEncoder(w).Encode(v)
}

func (app *App) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "all good")
}

// get the give user for the request out of the context and attempt to format
// the interface into the struct model
func GetUser(r *http.Request) (database.User, error) {
	return r.Context().Value("user").(database.User), nil
}
