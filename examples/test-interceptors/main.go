package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// Read and print the header
		log.Print("")
		log.Printf("Header: %v", request.Header)

		// Read and print the request body
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Printf("Error reading body: %s", err)
			http.Error(writer, "cannot read body", http.StatusBadRequest)
		}
		log.Printf("Body: %s", string(body))

		// Create response
		if _, err := writer.Write(body); err != nil {
			log.Printf("Error writing response: %s", err)
		}
		log.Printf("writer.Header(): %v", writer.Header())
	})

	log.Print("Listen and serve on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
