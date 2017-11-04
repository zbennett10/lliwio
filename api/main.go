package main

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

	return hex
}

func parseWords(request *http.Request) []string {
	var words WordPost
	err := json.NewDecoder(request.Body).Decode(&words)

	if err != nil {
		panic(err)
	}

	return words.Colors
}

func main() {
	http.HandleFunc("/", generateColors)

	log.Fatal(http.ListenAndServe(":8080", nil))
}