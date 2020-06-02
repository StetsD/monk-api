package helpers

import "net/http"

func Throw(w http.ResponseWriter, code int, customStatus *string) {
	var statusText string

	if *customStatus != "" {
		statusText = *customStatus
	} else {
		statusText = http.StatusText(code)
	}

	http.Error(w, statusText, code)
}
