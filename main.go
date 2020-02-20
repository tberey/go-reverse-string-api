package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	//Assign Port (8080) flag.
	portInput := flag.String("port", "8080", "Server init port")
	flag.Parse()

	//Set HTTP handlers
	http.HandleFunc("/reverse", reverseHandler)

	//Start HTTP server
	err := http.ListenAndServe(":"+*portInput, nil)
	if err != nil {
		fmt.Printf("http server shutdown: %s", err)
	}
}

//reverseHandler function is a HTTP handler which takes a "text" query string and returns it reversed.
func reverseHandler(w http.ResponseWriter, r *http.Request) {

	//GET method to assign any string found paired with the "text" key, in the query string.
	textQuery := r.URL.Query().Get("text")

	//Error handling: Write error code (402), if empty string returned from GET method performed on query string.
	if textQuery == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Call func to reverse text from GET method, and write to response body.
	w.Write([]byte(reverse(textQuery)))
}

// Pass text string into reverse function, to reverse character order of the it.
func reverse(text string) string {
	var revWord string
	for i := len(text) - 1; i > -1; i-- {
		revWord += string(text[i]) //Concatonate each character as string.
	}
	return revWord
}
