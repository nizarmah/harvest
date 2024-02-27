package auth

import (
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/driver/view"
)

var NewLogin = view.New[viewmodel.LoginViewData]
var NewSignup = view.New[viewmodel.SignUpViewData]
