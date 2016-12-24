package httpio

import "reflect"

type ApiErrors map[string]string

// ApiRequest interface to define an api request, it must implement
// a validate method which will check the fields are correct
// and return a bool representation of this
type ApiRequest interface {
	Validate() bool
	GetErrors() ApiErrors
}

//  Base request which can hold the errors used across
//  all api requests
type BaseRequest struct {
	Errors ApiErrors `json:"-"`
}

// GetErrors gets the errors of the Base Request in a more fluent fashion
// so that it complies with the interface
func (b *BaseRequest) GetErrors() ApiErrors {
	return b.Errors
}

// NotEmpty checks that a field is not empty
func (b *BaseRequest) NotEmpty(name string, value interface{}) bool {
	v := reflect.ValueOf(value)
	if v.Interface() == reflect.Zero(v.Type()).Interface() {
		b.Errors[name] = "The " + name + " field is required"
		return false
	}

	return true
}

// Empty checks that a field is empty
func (b *BaseRequest) Empty(name string, value interface{}) bool {
	return !b.NotEmpty(name, value)
}
