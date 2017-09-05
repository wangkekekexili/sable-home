package sable_home

import (
	"html/template"
	"net/http"
)

var (
	tributeTemplate          = template.Must(template.ParseFiles("tribute/index.go.html"))
	tributeBootstrapTemplate = template.Must(template.ParseFiles("tribute/bootstrap.go.html"))
)

func tribute(w http.ResponseWriter, r *http.Request) {
	tributeTemplate.Execute(w, nil)
}

func tributeBootstrap(w http.ResponseWriter, _ *http.Request) {
	tributeBootstrapTemplate.Execute(w, nil)
}
