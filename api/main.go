package main

import "fmt"
import "net/http"
import "log"
import "encoding/json"
import "encoding/hex"

type HexResponse struct {
	Colors []string
}

type WordPost struct {
	Colors []string `json:"colors"`
}

func generateColors(response http.ResponseWriter, request *http.Request) {
	words := parseWords(request)

	var hexResponse HexResponse
	for _, word := range words {
		hex := convertWordToHex(word)
		hexResponse.Colors = append(hexResponse.Colors, hex)
	}

	response.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(hexResponse.Colors)
	if err != nil {
		panic(err)
	}
	response.Write(jsonResponse)
}

func convertWordToHex(word string) string {
	hex := hex.EncodeToString([]byte(word))
	if len(hex) < 5 {
		for len(hex) < 5 {
			hex += "00"
		}
	}

	if len(hex) > 6 {
		hex = hex[0:6] 
	}

	return "#" + hex
}

func parseWords(request *http.Request) []string {
	var words WordPost
	err := json.NewDecoder(request.Body).Decode(&words)

	if err != nil {
		panic(err)
	}

	return words.Colors
}

func handleRoot(response http.ResponseWriter, request *http.Request) {
	switch(request.Method) {
	case "GET":
		fmt.Fprintf(response, "Success!")
	case "POST":
		generateColors(response, request)
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	fmt.Println("Starting lliwio api on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}