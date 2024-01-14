package helper

import (
	"groupie-tracker/api"
	"strconv"
	"strings"
)

// function to find matching data for first album
func AlbumDateFilter(start, end string, artistlist []string) []int {
	var list []int
	// convert start and end to integers
	strint, _ := strconv.Atoi(start)
	endint, _ := strconv.Atoi(end)
	for i, c := range artistlist {
		// split and take last index for year then convert it to int
		s := strings.Split(c, "-")
		sint, _ := strconv.Atoi(s[2])
		// if the year is equal or greater than start and equal or less than end add it
		if sint >= strint && sint <= endint {
			list = append(list, i+1)
		}
	}
	return list
}

// function to find matching data for Create date
func CreDatFilter(start, end string, artistlist map[int]int) []int {
	var list []int
	// convert start and end to integers
	strint, _ := strconv.Atoi(start)
	endint, _ := strconv.Atoi(end)
	for i, c := range artistlist {
		// if the year is equal or greater than start and equal or less than end add it
		if c >= strint && c <= endint {
			list = append(list, i)
		}
	}
	return list
}

// function to find matching data for member number
func MemberListFilter(NoList []string, artistlist []string) []int {
	var result [][]int
	// since there is 8 value for number of member add matching for every number selected
	for _, c := range NoList {
		if c != "" {
			result = append(result, NoMemeberFilter(c, artistlist))
		}
	}
	// Filter for all list givin for all numbers
	// use map to avoid duplicate
	Unique := make(map[int]bool)
	for _, l := range result {
		for _, v := range l {
			Unique[v] = true
		}
	}
	// add all id into one list
	var list []int
	for i := range Unique {
		list = append(list, i)
	}
	return list

}

// function to find matching number of memeber in the artist list givin
func NoMemeberFilter(n string, artistlist []string) []int {
	var list []int
	// convert the givin number into int
	num, _ := strconv.Atoi(n)
	for i, c := range artistlist {
		// convert the c number from string to int
		cnum, _ := strconv.Atoi(c)
		if cnum == num { // check if there is match
			list = append(list, i+1)
		}
	}
	return list
}

// function to find matching country location
func LocationFilter(name string, artistlist []api.Artists) []int {
	// use map to avoid duplicates
	var result []int
	for i, v := range artistlist {
		for _, c := range v.LOCATIONS {
			// the name should not be empty
			// check if the c (location) same as givin name
			if strings.ToLower(name) == c && name != "" {
				result = append(result, i+1)
			}
		}
	}
	return result
}
