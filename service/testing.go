package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper function to create a set of mock request response entities for testing
// specficically takes a slice of bytes json and creates a new httptest request
func NewJsonPostRequest(route string, json []byte) (*httptest.ResponseRecorder, *http.Request) {
	return newPutOrPostRequest("POST", route, json)
}

// helper function to create a set of mock request response entities for testing
// specficically takes a slice of bytes json and creates a new httptest request
func NewJsonPutRequest(route string, json []byte) (*httptest.ResponseRecorder, *http.Request) {
	return newPutOrPostRequest("PUT", route, json)
}

func newPutOrPostRequest(method, route string, json []byte) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()

	r := httptest.NewRequest(method, route, bytes.NewBuffer(json))
	r.Header.Set("Content-Type", "application/json")

	return w, r
}

// helper function to create a set of mock request response entities for testing
// creates a get request
func NewGetRequest(route string) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", route, nil)

	return w, r
}

// check that the body matches a given value
func SeeBodyMatches(t *testing.T, w *httptest.ResponseRecorder, expected string) {
	actual, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expected, strings.TrimSpace(string(actual)))
}
