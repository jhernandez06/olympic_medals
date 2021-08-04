package api_test

import (
	"net/http"
	"net/http/httptest"
	"olympic-medals/api"

	"testing"
)

func TestGetCountries(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/countries", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetCountries)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	} else {
		t.Log("StatusOK")
	}

}
