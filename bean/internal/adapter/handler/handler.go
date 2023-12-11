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
