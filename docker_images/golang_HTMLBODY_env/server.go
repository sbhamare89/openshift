package main

import (
	"net/http"
	"os"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	myVar := getEnv("HTMLBODY", "This text returned as value of env variable HTMLBODY from container")
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}
	w.Write([]byte("[ Hostname ] => " + hostname + ", [ Value of ENV variable HTMLBODY ] => " + myVar))
	return
}