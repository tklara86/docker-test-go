package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Docker tutorial")

		if err !=nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
