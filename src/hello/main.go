package main

import (
	"log"
	"net/http"
)

func main() {
	<break>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
