package helper

import (
	"sort"
	"strconv"
	"strings"
)

// struct for storing filter variables for the html
type FilterInputData struct {
	UniFirstAlbumData []string
	UniCreationData   [2]string
	UniMemberData     []string
	UniLocationData   []string
}

// function to add all variables into the struct
func AddFilterInputData(data FilterData) FilterInputData {
	InputData := FilterInputData{}
	var CreationDateList []string
	for _, v := range data.CreationData {
		CreationDateList = append(CreationDateList, strconv.Itoa(v))
	}
	sort.Strings(CreationDateList)
	InputData.UniCreationData[0] = CreationDateList[0]                       // add oldest date
	InputData.UniCreationData[1] = CreationDateList[len(CreationDateList)-1] // add latest date
	InputData.UniFirstAlbumData = UniqueData(data.FirstAlbumData)            // add first album list
	InputData.UniMemberData = UniqueData(data.MemberData)                    // add member list
	InputData.UniLocationData = UniqueDataMap(data.LocationData)             // add location
	return InputData
}

// function to avoid duplication
func UniqueData(data []string) []string {
	result := make(map[string]bool)
	for _, input := range data {
		if strings.Contains(input, "-") { // check for text if it is contain "-" or not
			arr := strings.Split(input, "-")
			input = arr[len(arr)-1]
		}
		// add the index and the value
		result[input] = true
	}
	var unique []string
	// loop to add index of result into list
	for i := range result {
		unique = append(unique, i)
	}
	// sort the list to arrange the array
	sort.Strings(unique)
	return unique
}

// function to add countries into the list
func UniqueDataMap(data map[string]int) []string {
	var countries []string
	for i := range data {
		// split the text and add last index as a country name
		countries = append(countries, strings.ToUpper(i))
	}
	result := make(map[string]bool)
	for _, input := range countries {
		// add the index and the value
		result[input] = true
	}
	var unique []string
	// loop to add index of result into list
	for i := range result {
		unique = append(unique, i)
	}
	// sort the list to arrange the array
	sort.Strings(unique)
	return unique
}
