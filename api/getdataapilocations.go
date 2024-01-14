package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// Locations struct locations.json
type Locations struct {
	ID        int      `json:"id"`
	LOCATIONS []string `json:"locations"`
	DATES     string   `json:"dates"`
}

func GetDataApiLocations(url string) (Locations, error) {
	// get the data from the link
	resp, err := http.Get(url)
	if err != nil {
		return Locations{}, err
	}
	// read the data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}
	var data Locations
	// parse the data into the data Locations
	json.Unmarshal(body, &data)
	return data, nil
}
