package api

import (
	"fmt"
	"strconv"
)

// select the artist by id number
func GetArtistData(artistID string, data []Artists) (Artists, error) {
	artid, err := strconv.Atoi(artistID)
	if artid > len(data) || err != nil {
		return Artists{}, fmt.Errorf("artist not found")
	}
	Artists := data[artid-1]

	return Artists, nil
}
