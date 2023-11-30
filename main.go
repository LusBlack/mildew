package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	//we set the method of requests
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "greetings")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
