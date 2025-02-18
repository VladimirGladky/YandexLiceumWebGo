package main

import (
	"errors"
	"fmt"
	"net/http"
)

func FourthTaskHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		http.SetCookie(w, &http.Cookie{
			Name:  "session_id",
			Value: "abc123",
		})
		fmt.Fprintln(w, "Welcome!")
		return
	}
	fmt.Fprintln(w, "Welcome back!")
}

func StartServer4() {
	http.HandleFunc("/", FourthTaskHandler)
	http.ListenAndServe(":8080", nil)
}
