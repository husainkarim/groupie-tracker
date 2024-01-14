package server

import (
	"html/template"
	"net/http"
)

type Error struct {
	Title string
	Body  string
}

// function to hundle 400 error response
func ErrorHandler(w http.ResponseWriter, r *http.Request, n int) {
	index := "templates/error.html"
	tmpl, err := template.ParseFiles(index) // make template for the error page
	if err != nil {                         // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	data := &Error{}
	if n == 400 { // 400 error response
		data.Title = "Bad Request 400"
		data.Body = "Bad Request 400"
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
	}
	if n == 404 { // 404 error response
		data.Title = "Page Not Found 404"
		data.Body = "Page Not Found 404"
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
	}
	if n == 500 { // 500 error response
		data.Title = "Internal Server Error 500"
		data.Body = "Internal Server Error 500"
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err1 := tmpl.Execute(w, data) // execute the template wiht the result
	if err1 != nil {              // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
