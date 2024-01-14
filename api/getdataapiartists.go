package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// Artists struct match artists.json var type and logic
type Artists struct {
	ID           int                 `json:"id"`
	IMAGE        string              `json:"image"`
	NAME         string              `json:"name"`
	MEMBERS      []string            `json:"members"`
	CREA_DATE    int                 `json:"creationDate"`
	FIRST_ALBUM  string              `json:"firstAlbum"`
	LOCATIONS    []string            `json:"locations"`
	CONCERT_DATE []string            `json:"concertDates"`
	RELATION     map[string][]string `json:"relations"`
}

func GetDataApiArtists(url string) ([]Artists, error) {
	// get the data from the link
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// read the data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data []Artists
	// parse the data into the data []Artists
	json.Unmarshal(body, &data)
	return data, nil
}
