package main

import (
	"net/http"
)

func generateColors(response http.ResponseWriter, request *http.Request) {
		
}

func main() {
	http.HandleFunc("/", generateColors)
}