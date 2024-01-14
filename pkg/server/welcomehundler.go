package server

import (
	"html/template"
	"net/http"
)

func WelcomeHundler(w http.ResponseWriter, r *http.Request) {
	index := "templates/welcome.html"
	tmpl, err := template.ParseFiles(index) // make template for the page
	if err != nil {                         // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	err1 := tmpl.Execute(w, nil) // execute the template wiht the result
	if err1 != nil {             // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
