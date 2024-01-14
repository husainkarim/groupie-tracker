package server

import (
	"groupie-tracker/api"
	"groupie-tracker/pkg/helper"
	"html/template"
	"net/http"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request, data ServerData) {
	var artistData api.Artists
	var err error
	if r.URL.Query().Get("id") != "" {
		artistID := r.URL.Query().Get("id")                            //get the id from the path
		artistData, err = api.GetArtistData(artistID, data.ArtistData) // get the data for the artist selected
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			return
		}
	} else { // if no id provided then go to list page
		ArtistEmptyHandler(w, r, data)
		return
	}
	layout := "templates/layout.html"
	index := "templates/artist.html"
	tmpl, err2 := template.ParseFiles(layout, index) // make template for the page
	if err2 != nil {                                 // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	SortedLocation := helper.LocationSort(artistData)
	artistData.LOCATIONS = SortedLocation
	err1 := tmpl.Execute(w, artistData) // execute the template with the result
	if err1 != nil {                    // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
