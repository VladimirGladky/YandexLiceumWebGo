package main

import (
	"fmt"
	"net/http"
)

func SeventhTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authorized access")
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer valid_token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
