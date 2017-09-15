package sable_home

import (
	"html/template"
	"net/http"
)

type headFirst struct {
	starbuzz http.Handler
	lounge   http.Handler
}

func newHeadFirst() *headFirst {
	return &headFirst{
		starbuzz: newStarBuzz(),
		lounge:   newLounge(),
	}
}

func (h *headFirst) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current string
	current, r.URL.Path = shiftPath(r.URL.Path)
	switch current {
	case "starbuzz":
		h.starbuzz.ServeHTTP(w, r)
	case "lounge":
		h.lounge.ServeHTTP(w, r)
	}
}

type starbuzz struct {
	index   *template.Template
	mission *template.Template
}

func newStarBuzz() *starbuzz {
	return &starbuzz{
		index:   template.Must(template.ParseFiles("headfirst/starbuzz/index.go.html")),
		mission: template.Must(template.ParseFiles("headfirst/starbuzz/mission.go.html")),
	}
}

func (s *starbuzz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	current, _ := shiftPath(r.URL.Path)
	switch current {
	case "index":
		s.index.Execute(w, nil)
	case "mission":
		s.mission.Execute(w, nil)
	}
}

type lounge struct {
	index      *template.Template
	elixir     *template.Template
	directions *template.Template
}

func newLounge() *lounge {
	return &lounge{
		index:      template.Must(template.ParseFiles("headfirst/lounge/lounge.html")),
		elixir:     template.Must(template.ParseFiles("headfirst/lounge/elixir.html")),
		directions: template.Must(template.ParseFiles("headfirst/lounge/directions.html")),
	}
}

func (g *lounge) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	current, _ := shiftPath(r.URL.Path)
	switch current {
	case "index":
		g.index.Execute(w, nil)
	case "elixir":
		g.elixir.Execute(w, nil)
	case "directions":
		g.directions.Execute(w, nil)
	}
}
