package server

import (
	"groupie-tracker/api"
	"groupie-tracker/pkg/helper"
	"net/http"
	"sort"
)

// main variable to store all requirement and handle all requests from the users
type ServerData struct {
	ArtistData      []api.Artists
	SearchData      []string
	FilterData      helper.FilterData
	FirstYear       string
	LastYear        string
	StartAlbumYear  string
	EndAlbumYear    string
	UniMemberData   []string
	UniLocationData []string
	AlartFilter     string
	AlartSearch     string
	CurrentCity     string
}

// all data will be stored in this global variable
var data ServerData

func init() {
	var err error
	var w http.ResponseWriter
	var r *http.Request
	// add artist data list tp global variable
	data.ArtistData, err = api.ApiData()
	if err != nil { // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
	// add searching data and filter data to the global variable
	data.SearchData = helper.SearchBarData(data.ArtistData)
	data.FilterData = helper.AddFilterData(data.ArtistData)
	var FilterInputData helper.FilterInputData = helper.AddFilterInputData(data.FilterData)
	data.FirstYear = FilterInputData.UniCreationData[0]
	data.LastYear = FilterInputData.UniCreationData[1]
	AlbumYear := FilterInputData.UniFirstAlbumData
	sort.Strings(AlbumYear)
	data.StartAlbumYear = AlbumYear[0]
	data.EndAlbumYear = AlbumYear[len(AlbumYear)-1]
	data.UniMemberData = FilterInputData.UniMemberData
	data.UniLocationData = FilterInputData.UniLocationData
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": // home page
		MainHandler(w, r, data)
	case "/artist": // artist page
		ArtistHandler(w, r, data)
	case "/about": // about page
		AboutHandler(w, r)
	case "/welcome": // welcome page
		WelcomeHundler(w, r)
	case "/filter": // filter page
		FilterHundler(w, r, data)
	case "/search": // search page
		SearchHundler(w, r, data)
	case "/400":
		ErrorHandler(w, r, 400)
	case "/500":
		ErrorHandler(w, r, 500)
	default: // any invalid path will be redirected to 404 error handler
		ErrorHandler(w, r, 404)
	}
}
