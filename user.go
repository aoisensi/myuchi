package main

import "net/http"

func userCreate(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "POST") {
		return
	}
}
