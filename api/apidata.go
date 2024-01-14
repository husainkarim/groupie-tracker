package api

import (
	"fmt"
)

// get the remaining empty data to add it to the Artists list
func ApiData() ([]Artists, error) {
	artists, err := GetDataApiArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	results := make(chan int, len(artists))
	for i := range artists {
		go func(i int) error {
			defer func() { results <- i }()
			// add locations
			locations, err := GetDataApiLocations("https://groupietrackers.herokuapp.com/api/locations/" + fmt.Sprint(i+1))
			if err != nil {
				return err
			}
			artists[i].LOCATIONS = locations.LOCATIONS
			// add dates
			dates, err := GetDataApiDates("https://groupietrackers.herokuapp.com/api/dates/" + fmt.Sprint(i+1))
			if err != nil {
				return err
			}
			artists[i].CONCERT_DATE = dates.DATES
			// add relation
			relation, err := GetDataApiRelations("https://groupietrackers.herokuapp.com/api/relation/" + fmt.Sprint(i+1))
			if err != nil {
				return err
			}
			artists[i].RELATION = relation.DATESLOCAT
			return nil
		}(i)
	}
	for range artists {
		<-results
	}
	return artists, nil
}
