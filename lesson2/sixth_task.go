package main

import (
	"fmt"
	"net/http"
)

func SixthTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Access granted")
}

func ipBlockerMiddleware(blockedIP string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Real-IP") == blockedIP {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
