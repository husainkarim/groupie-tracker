package server

import (
	"groupie-tracker/api"
	"groupie-tracker/pkg/helper"
	"net/http"
	"strconv"
)

func SearchHundler(w http.ResponseWriter, r *http.Request, data ServerData) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var MainData ServerData = data
	var SearchList []api.Artists
	search := r.PostFormValue("search") // get the post search text
	if search != "" {
		idList := helper.Search(search, data.ArtistData) // identify the artist id
		for _, id := range idList {                      // add the artists to the search list created for the search
			arty, _ := api.GetArtistData(strconv.Itoa(id), data.ArtistData)
			SearchList = append(SearchList, arty)
		}
		if SearchList != nil {
			MainData.ArtistData = SearchList
		} else {
			MainData.AlartSearch = "No Matching Search"
		}
	}
	// redirect to main page to display the search results
	MainHandler(w, r, MainData)
}
