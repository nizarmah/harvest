package app

import (
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/driver/view"
)

var NewHome = view.New[viewmodel.HomeViewData]
var NewOnboarding = view.New[viewmodel.OnboardingViewData]
