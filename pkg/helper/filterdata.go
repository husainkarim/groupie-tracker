package helper

import (
	"groupie-tracker/api"
	"strconv"
)

type FilterData struct {
	FirstAlbumData []string
	CreationData   map[int]int
	MemberData     []string
	LocationData   map[string]int
}

// function to add all filter data into one struct
func AddFilterData(data []api.Artists) FilterData {
	filterdata := FilterData{}
	filterdata.FirstAlbumData = ArtistFirstAlbum(data)
	filterdata.CreationData = ArtistCreationDate(data)
	filterdata.MemberData = ArtistMembers(data)
	filterdata.LocationData = ArtistLocations(data)
	return filterdata

}

func ArtistFirstAlbum(data []api.Artists) []string {
	var result []string // use map variable to avoid duplicates of elements
	for _, v := range data {
		result = append(result, v.FIRST_ALBUM) // add first album date
	}
	return result
}

func ArtistCreationDate(data []api.Artists) map[int]int {
	result := make(map[int]int) // use map variable to avoid duplicates of elements
	for i, v := range data {
		result[i+1] = v.CREA_DATE // add creation date
	}
	return result
}

func ArtistMembers(data []api.Artists) []string {
	var result []string // use map variable to avoid duplicates of elements
	for _, v := range data {
		result = append(result, strconv.Itoa(len(v.MEMBERS))) // add number of members
	}
	return result
}

func ArtistLocations(data []api.Artists) map[string]int {
	result := make(map[string]int) // use map variable to avoid duplicates of elements
	for i, v := range data {
		for _, l := range v.LOCATIONS { // add locations
			result[l] = i + 1
		}
	}
	return result
}
