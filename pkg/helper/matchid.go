package helper

func MatchId(CreatDateList, FirstAlbumList, MemberList, LocationList []int) []int {
	var matchlist []int
	// list to add all members id to empty list
	ArtList := AddArtist()
	if CreatDateList == nil { // add all members id to create date list if it is empty
		CreatDateList = ArtList
	}
	if FirstAlbumList == nil { // add all members id to first album date list if it is empty
		FirstAlbumList = ArtList
	}
	if MemberList == nil { // add all members id to member number list if it is empty
		MemberList = ArtList
	}
	if LocationList == nil { // add all members id to location list if it is empty
		LocationList = ArtList
	}
	// loop to check matching members to add it ton the list
	for i := 1; i <= 52; i++ {
		if ContainList(CreatDateList, i) && ContainList(FirstAlbumList, i) && ContainList(MemberList, i) && ContainList(LocationList, i) {
			matchlist = append(matchlist, i)
		}
	}
	return matchlist
}

func ContainList(list []int, n int) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}

func AddArtist() []int {
	var res []int
	for i := 1; i <= 52; i++ {
		res = append(res, i)
	}
	return res
}
