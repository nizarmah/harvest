package template

import (
	"embed"
)

//go:embed *.html layout/*.html
//go:embed paymentmethod/*.html
var FS embed.FS

const (
	baseTemplate = "layout/base.html"
)

var (
	LandingTemplate = []string{
		baseTemplate,
		"landing.html",
	}

	LoginTemplate = []string{
		baseTemplate,
		"login.html",
	}

	HomeTemplate = []string{
		baseTemplate,
		"home.html",
	}

	CreatePaymentMethodTemplate = []string{
		baseTemplate,
		"paymentmethod/create.html",
	}
)
