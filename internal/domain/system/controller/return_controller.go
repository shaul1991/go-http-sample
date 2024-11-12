package controller

import (
	"fmt"
	"net/http"
)

// ReturnMessage responds with a 200 status and "hello world" message
func ReturnMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "hello world")
}
