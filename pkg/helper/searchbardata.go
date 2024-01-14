package helper

import (
	"groupie-tracker/api"
	"sort"
	"strconv"
)

func SearchBarData(data []api.Artists) []string {
	result := make(map[string]bool) // use map variable to avoid duplicates of elements
	for _, v := range data {
		result[v.NAME+" - Artist/Band"] = true                      // add artist name
		result[v.FIRST_ALBUM+" - 1st Album Date"] = true            // add first album date
		result[strconv.Itoa(v.CREA_DATE)+" - Creation Date"] = true // add creation date
		for _, m := range v.MEMBERS {                               // add members
			result[m+" - Member"] = true
		}
		for _, l := range v.LOCATIONS { // add locations
			result[l+" - Location"] = true
		}
	}
	var list []string
	for i := range result { // add index of the map only
		list = append(list, i)
	}
	sort.Strings(list) // sort results
	return list
}
