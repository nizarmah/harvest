package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Landing(w http.ResponseWriter, r *http.Request) {
	e := render(w, "landing.html", nil)
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")

	var errmsg string
	if r.Method == "POST" {
		errmsg = fmt.Sprintf("Account '%v' not found", username)
	}

	e := render(w, "login.html", map[string]string{
		"Error": errmsg,
	})
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}
}

func render(w http.ResponseWriter, t string, data interface{}) error {
	lp := filepath.Join("./static/template", "_layout.html")
	tp := filepath.Join("./static/template", t)

	tmpl, e := template.ParseFiles(lp, tp)
	if e != nil {
		return e
	}

	e = tmpl.Execute(w, data)
	if e != nil {
		return e
	}

	return nil
}
