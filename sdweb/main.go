package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "URL.Path = %q\n", req.URL)
}
