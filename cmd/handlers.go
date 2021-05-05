package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	comms, err := app.comments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	file := "/home/prokiprok/go/src/kursach2.0/html/home_page.html"

	// ts, err := template.ParseFiles(file)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		app.errorLog.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "comments", comms)
	// err = tmpl.Execute(w, nil)
	// err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach20/html/about.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) sources(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sources" {
		app.notFound(w)
		return
	}

	files := []string{
		"/home/prokiprok/go/src/kursach20/html/sources.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
	}
}
