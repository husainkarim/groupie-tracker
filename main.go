package main

import (
	"fmt"
	"groupie-tracker/pkg/server"
	"net/http"
)

func main() {
	// path handler function to handle all possible paths
	http.HandleFunc("/", server.PathHandler)
	// local host link
	link := "http://localhost:8080/welcome"
	fmt.Println("\033[36mServer Connected...\033[0m")
	fmt.Printf("\033[36mlink on: %s\033[0m\n", link)
	// Connect the css & img & js & font folder to the server
	http.Handle("/css/", http.FileServer(http.Dir("templates")))
	http.Handle("/img/", http.FileServer(http.Dir("templates")))
	http.Handle("/js/", http.FileServer(http.Dir("templates")))
	http.Handle("/fonts/", http.FileServer(http.Dir("templates")))
	// running the server will be on port 8080.
	http.ListenAndServe(":8080", nil)
}
