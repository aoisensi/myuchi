package main

import (
	"fmt"
	"net/http"
)

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		fmt.Fprintf(w, "Not %s method", method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
