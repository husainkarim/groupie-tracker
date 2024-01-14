package server

import (
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request, data ServerData) {
	layout := "templates/layout.html"
	index := "templates/index.html"
	tmpl, err := template.ParseFiles(layout, index) // make template for the page
	if err != nil {                                 // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	err1 := tmpl.Execute(w, data) // execute the template with the result
	if err1 != nil {              // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
