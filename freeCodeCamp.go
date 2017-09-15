package sable_home

import (
	"html/template"
	"net/http"
)

type freeCodeCamp struct {
	tribute          *template.Template
	tributeBootstrap *template.Template
}

func newFreeCodeCamp() *freeCodeCamp {
	return &freeCodeCamp{
		tribute:          template.Must(template.ParseFiles("freecodecamp/tribute/index.go.html")),
		tributeBootstrap: template.Must(template.ParseFiles("freecodecamp/tribute/bootstrap.go.html")),
	}
}

func (f *freeCodeCamp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head, _ := shiftPath(r.URL.Path)
	switch head {
	case "tribute":
		f.tribute.Execute(w, nil)
	case "tribute-bootstrap":
		f.tributeBootstrap.Execute(w, nil)
	}
}
