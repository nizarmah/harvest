package marketing

import (
	"fmt"
	"net/http"
)

func (h *Handlers) Landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Bean!")
}
