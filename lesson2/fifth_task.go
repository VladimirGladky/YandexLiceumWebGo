package main

import (
	"errors"
	"fmt"
	"net/http"
)

func FifthTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Access granted")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: "abc123",
	})
	fmt.Fprintln(w, "Please log in")
}

func FifthTaskMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("user_id")
		if errors.Is(err, http.ErrNoCookie) {
			http.Redirect(w, r, "/login", 301)
			return
		}
		next(w, r)
	}
}

func StartServer5() {
	http.HandleFunc("/", FifthTaskMW(FifthTaskHandler))
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8080", nil)
}
