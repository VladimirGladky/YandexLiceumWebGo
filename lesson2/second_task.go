package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SecondTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server")
}

func sendRequest(url string) (string, error) {
	if strings.HasPrefix(url, "localhost") {
		url = strings.Replace(url, "localhost", "http://127.0.0.1", 1)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf(" sendRequest,NewRequest : %v", err)
	}
	resp, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		return "", fmt.Errorf(" sendRequest,DefaultClientDo : %v", err1)
	}
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return "", fmt.Errorf(" sendRequest,ReadAll : %v", err2)
	}
	return string(body), nil
}

func startServer(address string) {
	http.HandleFunc("/", SecondTaskHandler)
	http.ListenAndServe(address, nil)
}
