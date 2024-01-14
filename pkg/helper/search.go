package helper

import (
	"groupie-tracker/api"
	"sort"
	"strconv"
	"strings"
)

func Search(search string, data []api.Artists) []int {
	result := make(map[int]bool) // use map variable to avoid duplicates of elements
	var searchelement string
	var area string
	if strings.Contains(search, " - ") { // if there is " - " that means the text needs to split into 2 parts
		searchlist := strings.Split(search, " - ")
		searchelement = strings.ToLower(searchlist[0]) // take the first part to check the matching text in the text list
		area = searchlist[1]
	} else { // if there is no " - " take the full text as search
		searchelement = strings.ToLower(search)
	}
	for i, a := range data {
		if area == "" {
			if strings.Contains(strings.ToLower(a.NAME), searchelement) { // check the matching in artist name
				result[i+1] = true
			}
			if strings.Contains(strings.ToLower(a.FIRST_ALBUM), searchelement) { // check the matching in first album date
				result[i+1] = true
			}
			if strings.Contains(strconv.Itoa(a.CREA_DATE), searchelement) { // check the matching in creation date
				result[i+1] = true
			}
			for _, m := range a.MEMBERS { // check the matching in members name
				if strings.Contains(strings.ToLower(m), searchelement) {
					result[i+1] = true
				}
			}
			for _, l := range a.LOCATIONS { // check the matching in locations name
				if strings.Contains(strings.ToLower(l), searchelement) {
					result[i+1] = true
				}
			}
		} else {
			switch area {
			case "Artist/Band":
				if strings.Contains(strings.ToLower(a.NAME), searchelement) { // check the matching in artist name
					result[i+1] = true
				}
			case "1st Album Date":
				if strings.Contains(strings.ToLower(a.FIRST_ALBUM), searchelement) { // check the matching in first album date
					result[i+1] = true
				}
			case "Creation Date":
				if strings.Contains(strconv.Itoa(a.CREA_DATE), searchelement) { // check the matching in creation date
					result[i+1] = true
				}
			case "Member":
				for _, m := range a.MEMBERS { // check the matching in members name
					if strings.Contains(strings.ToLower(m), searchelement) {
						result[i+1] = true
					}
				}
			case "Location":
				for _, l := range a.LOCATIONS { // check the matching in locations name
					if strings.Contains(strings.ToLower(l), searchelement) {
						result[i+1] = true
					}
				}
			}
		}
	}
	var list []int
	for l := range result { // add index of the map only
		list = append(list, l)
	}
	sort.Ints(list) // sort the list
	return list
}
