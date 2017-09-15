package sable_home

import (
	"net/http"
	"path/filepath"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", newApp())
}

type app struct {
	freeCodeCamp http.Handler
	headFirst    http.Handler
}

func newApp() *app {
	return &app{
		freeCodeCamp: newFreeCodeCamp(),
		headFirst:    newHeadFirst(),
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var current string
	log.Infof(appengine.NewContext(r), r.URL.Path)
	current, r.URL.Path = shiftPath(r.URL.Path)
	log.Infof(appengine.NewContext(r), current)
	switch current {
	case "freecodecamp":
		a.freeCodeCamp.ServeHTTP(w, r)
	case "headfirst":
		a.headFirst.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func shiftPath(p string) (head, tail string) {
	p = filepath.Clean("/" + p)
	splitIndex := strings.Index(p[1:], "/")
	if splitIndex < 0 {
		return p[1:], "/"
	}
	return p[1:splitIndex+1], p[splitIndex+1:]
}
