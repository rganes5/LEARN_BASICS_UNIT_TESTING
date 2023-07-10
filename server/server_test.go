package server

import (
	"net/http/httptest"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {

}
