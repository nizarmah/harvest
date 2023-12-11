package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Landing(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("./static/template/landing.html")
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}

	e = t.Execute(w, nil)
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	t, e := template.ParseFiles("./static/template/login.html")
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}

	var errmsg string
	if r.Method == "POST" {
		errmsg = fmt.Sprintf("Account '%v' not found", username)
	}

	data := struct {
		Error string
	}{
		Error: errmsg,
	}

	e = t.Execute(w, data)
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}
}
