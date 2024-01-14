package helper

import (
	"groupie-tracker/api"
	"log"
	"sort"
	"time"
)

type LocationSorted struct {
	key  string
	date time.Time
}

func LocationSort(data api.Artists) []string {
	// Create a new slice to store the sorted date locations.
	sortedDateLocations := []LocationSorted{}

	// Iterate over the map and for each key-value pair:
	// * Convert the date string to a `time.Time` object.
	// * Append the key-value pair to the slice, sorted by date.
	for key, value := range data.RELATION {
		date, err := time.Parse("02-01-2006", value[0])
		if err != nil {
			log.Fatal(err)
		}
		dateLocation := LocationSorted{key: key, date: date}
		sortedDateLocations = append(sortedDateLocations, dateLocation)
	}

	// Sort the slice of date locations by date.
	sort.Slice(sortedDateLocations, func(i, j int) bool {
		return !sortedDateLocations[i].date.Before(sortedDateLocations[j].date)
	})

	// Create a new slice to store the sorted keys.
	sortedKeys := []string{}

	// Iterate over the slice of sorted date locations and append the keys to the slice of sorted keys.
	for _, dateLocation := range sortedDateLocations {
		sortedKeys = append(sortedKeys, dateLocation.key)
	}
	return sortedKeys
}
