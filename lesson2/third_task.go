package main

import (
	"fmt"
	"log"
	"net/http"
)

func ThirdTaskHandler(w http.ResponseWriter, r *http.Request) {
	l := log.Logger{}
	l.Println(r.Method)
	l.Println(r.URL)
}

func ThirdTaskMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "Middleware Test")
		}
		next(w, r)
	}
}

func StartServer3() {
	http.HandleFunc("/", ThirdTaskMW(ThirdTaskHandler))
	http.ListenAndServe(":8080", nil)
}
