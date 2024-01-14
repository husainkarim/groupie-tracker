package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// Dates struct match dates.json var type and logic
type Dates struct {
	ID    int      `json:"id"`
	DATES []string `json:"dates"`
}

func GetDataApiDates(url string) (Dates, error) {
	// get the data from the link
	resp, err := http.Get(url)
	if err != nil {
		return Dates{}, err
	}
	// read the data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Dates{}, err
	}
	var data Dates
	// parse the data into the data Dates
	json.Unmarshal(body, &data)
	return data, nil
}
