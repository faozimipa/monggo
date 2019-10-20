package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faozimipa/monggo/api"
)

func TestUrlInAPI(t *testing.T) {
	t.Run("check home url", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/entries", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(api.GetAllStudents)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := 1
		if len(rr.Body.String()) <= expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}

	})
}
