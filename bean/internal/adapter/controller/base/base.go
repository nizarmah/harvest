package base

import (
	"fmt"
	"net/http"
)

type Controller struct{}

func (c *Controller) ErrorHandler(
	handler HTTPHandler,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		switch e := err.(type) {
		case *HTTPError:
			handleHTTPError(w, r, e)
		default:
			handleGenericError(w, r, e)
		}
	}
}

func handleHTTPError(w http.ResponseWriter, r *http.Request, err *HTTPError) {
	if err == nil {
		return
	}

	fmt.Println("server: http error:", err)

	if err.Status == 0 {
		return
	}

	if err.Status >= http.StatusInternalServerError {
		http.Error(
			w,
			"Something went wrong",
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
