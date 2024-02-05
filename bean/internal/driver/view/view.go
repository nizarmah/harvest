package view

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/adapter/interfaces"
)

type Base[T entity.ViewData] struct {
	tmpl *template.Template
}

func New[T entity.ViewData](FS embed.FS, templates []string) (interfaces.View[T], error) {
	tmpl, err := template.ParseFS(FS, templates...)
	if err != nil {
		return nil, fmt.Errorf("error parsing templates: %w", err)
	}

	return &Base[T]{tmpl: tmpl}, nil
}

func (v *Base[T]) Render(w http.ResponseWriter, data *T) error {
	return v.tmpl.Execute(w, data)
}
