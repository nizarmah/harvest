package handler

import (
	"fmt"
	"html/template"
	"net/http"

	// FIXME: wrong direction of dependency
	templateDriver "harvest/bean/internal/driver/template"
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
	tmpl, e := template.ParseFS(templateDriver.FS, "_layout.html", t)
	if e != nil {
		return e
	}

	e = tmpl.Execute(w, data)
	if e != nil {
		return e
	}

	return nil
}
