package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, (r.URL.Path))
		fmt.Println(r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
