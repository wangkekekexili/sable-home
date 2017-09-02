package sable_home

import (
	"html/template"
	"net/http"
)

var tributeTemplate = template.Must(template.ParseFiles("tribute/index.go.html"))

func tribute(w http.ResponseWriter, r *http.Request) {
	tributeTemplate.Execute(w, nil)
}
