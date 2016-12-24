package httpio

import (
	"encoding/json"
	"errors"
	"net/http"
)

var UserOutOfContext = errors.New("User out of context")

// ErrorResponse is a helper that writes a error to a ResponseWriter setting the
// header to status code 400
func ErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(ErrorOutput{err.Error()})
}

// NotFound is a helper that writes a error to a ResponseWriter setting the
// header to status code 404
func NotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(ErrorOutput{message})
}

// WriteValidationResponse is a helper that writes a error to a ResponseWriter setting the
// header to status code 422
func WriteValidationResponse(w http.ResponseWriter, v *ValidationErrorResponse) {
	w.WriteHeader(422)
	json.NewEncoder(w).Encode(v)
}
