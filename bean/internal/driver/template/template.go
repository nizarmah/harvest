package template

import (
	"embed"
)

//go:embed layout/*.html *.html
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

	SubscriptionsTemplate = []string{
		baseTemplate,
		"subscriptions.html",
	}
)
