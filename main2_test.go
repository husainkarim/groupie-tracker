package main

import (
	"groupie-tracker/api"
	"testing"
)

func TestApiData(t *testing.T) {
	expected, err := api.ApiData()
	empty := []api.Artists{}
	if &expected == &empty && err == nil {
		t.Error("Error: there is problem in ApiData function!\n")
	}
}
