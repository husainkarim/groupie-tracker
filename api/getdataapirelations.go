package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// Relations struct match relations.json var type and logic
type Relations struct {
	ID         int                 `json:"id"`
	DATESLOCAT map[string][]string `json:"datesLocations"`
}

func GetDataApiRelations(url string) (Relations, error) {
	// get the data from the link
	resp, err := http.Get(url)
	if err != nil {
		return Relations{}, err
	}
	// read the data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Relations{}, err
	}
	var data Relations
	// parse the data into the data Relations
	json.Unmarshal(body, &data)
	return data, nil
}
