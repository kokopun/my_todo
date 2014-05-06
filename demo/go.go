package main

import (
    "fmt"
    "net/http"
)

func main() {
	 http.HandleFunc("/test", myAddHanler)
    http.HandleFunc("/", addHanler)
    http.ListenAndServe(":8080", nil)
}

func addHanler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
}

func myAddHanler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "test")
}



