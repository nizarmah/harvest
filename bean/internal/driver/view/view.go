package view

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/adapter/interfaces"
)

type Base[T viewmodel.ViewData] struct {
	tmpl *template.Template
}

func New[T viewmodel.ViewData](FS embed.FS, templates []string) (interfaces.View[T], error) {
	tmpl, err := template.ParseFS(FS, templates...)
	if err != nil {
		return nil, fmt.Errorf("error parsing templates: %w", err)
	}

	return &Base[T]{tmpl: tmpl}, nil
}

func (v *Base[T]) Render(w http.ResponseWriter, data *T) error {
	return v.tmpl.Execute(w, data)
}
