package server

import (
	"groupie-tracker/api"
	"groupie-tracker/pkg/helper"
	"net/http"
	"strconv"
)

func FilterHundler(w http.ResponseWriter, r *http.Request, data ServerData) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var MainData ServerData = data
	var FilterList []api.Artists
	year_start := r.PostFormValue("yearstart")     // get the start year for filtering creation year
	year_end := r.PostFormValue("yearend")         // get the end year for filtering creation year
	AlbumStart := r.PostFormValue("1stAlbumStart") // get the start year for first album date
	Albumend := r.PostFormValue("1stAlbumEnd")     // get the end year for first album date
	// get the number of the members for filtering
	no1 := r.PostFormValue("1")
	no2 := r.PostFormValue("2")
	no3 := r.PostFormValue("3")
	no4 := r.PostFormValue("4")
	no5 := r.PostFormValue("5")
	no6 := r.PostFormValue("6")
	no7 := r.PostFormValue("7")
	no8 := r.PostFormValue("8")
	country := r.PostFormValue("countries") // get the country name for filtering
	var NoMember []string
	if year_start != "" || year_end != "" || AlbumStart != "" || Albumend != "" || no1 != "" || no2 != "" || no3 != "" || no4 != "" || no5 != "" || no6 != "" || no7 != "" || no8 != "" || country != "" {
		if year_start > year_end || AlbumStart > Albumend { // Error if start date less than the end date
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}
		NoMember = append(NoMember, no1, no2, no3, no4, no5, no6, no7, no8)                        // add all members var into one list
		CreateDate := helper.CreDatFilter(year_start, year_end, data.FilterData.CreationData)      // get the matching list for creation date
		FirstAlbum := helper.AlbumDateFilter(AlbumStart, Albumend, data.FilterData.FirstAlbumData) // get the matching list for first album date
		Member := helper.MemberListFilter(NoMember, data.FilterData.MemberData)                    // get the matching list for member number
		cou := helper.LocationFilter(country, data.ArtistData)                                     // get the matching list for country
		matchId := helper.MatchId(CreateDate, FirstAlbum, Member, cou)                             // check matching list between all list
		for _, id := range matchId {                                                               // add the artists to the search list created for the search
			artx, _ := api.GetArtistData(strconv.Itoa(id), data.ArtistData)
			FilterList = append(FilterList, artx)
		}
		if FilterList != nil {
			MainData.ArtistData = FilterList // return the list of artist data matching the filter
		} else {
			MainData.AlartFilter = "No Matching Filter"
		}
	}
	// redirect to main page to display the filter results
	MainHandler(w, r, MainData)
}
