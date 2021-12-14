package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("launching server at port 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
