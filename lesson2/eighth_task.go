package main

import (
	"errors"
	"fmt"
	"net/http"
)

func languageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lang")
	if errors.Is(err, http.ErrNoCookie) {
		fmt.Fprintln(w, "Hello!")
		return
	}
	if cookie.Value == "ru" {
		fmt.Fprintln(w, "Привет!")
		return
	}
	fmt.Fprintln(w, "Hello!")
}
