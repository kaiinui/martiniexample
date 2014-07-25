package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"html/template"
)

func main() {
	m := martini.Classic()

	m.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		//html, _ := ioutil.ReadFile("views/index.html")
		name := r.FormValue("name")
		t, _ := template.ParseFiles("views/index.tmpl")
		t.Execute(w, name)
	})

	m.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("views/new.tmpl")
		t.Execute(w, nil)
	})

	m.Run()
}