package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNextSalaryHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/next-salary?payday=12", nil)
	if err != nil {
		t.Error("request failed ", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(nextSalaryHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("actual %v expected %v",
			status, http.StatusOK)
	}
}

func TestAnnualScheduleHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/annual_schedule?payday=12", nil)
	if err != nil {
		t.Error("request failed ", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(nextSalaryHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("actual %v expected %v",
			status, http.StatusOK)
	}
}
