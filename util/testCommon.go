package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

func ExecuteRequest(req *http.Request, ro *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ro.ServeHTTP(rr, req)
	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}