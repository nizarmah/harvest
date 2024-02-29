package base

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

type Controller struct{}

func (c *Controller) ErrorHandler(
	handler model.HTTPHandler,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		switch e := err.(type) {
		case *model.HTTPError:
			handleHTTPError(w, r, e)
		default:
			handleGenericError(w, r, e)
		}
	}
}

func handleHTTPError(w http.ResponseWriter, r *http.Request, err *model.HTTPError) {
	if err == nil {
		return
	}

	fmt.Println("server: http error:", err)

	if err.Status == 0 || err.Status >= http.StatusInternalServerError {
		http.Error(
			w,
			"Something went wrong",
			http.StatusInternalServerError,
		)
		return
	}

	if err.RedirectPath != "" {
		http.Redirect(
			w,
			r,
			err.RedirectPath,
			err.Status,
		)
		return
	}

	http.Error(
		w,
		err.Message,
		err.Status,
	)
}

func handleGenericError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	fmt.Println("server: generic error:", err)

	http.Error(
		w,
		"Something went wrong",
		http.StatusInternalServerError,
	)
}
