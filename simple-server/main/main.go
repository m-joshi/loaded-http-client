package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/simpleServer", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

