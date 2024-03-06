package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message struct represents the structure of the JSON response
type Message struct {
	Text string `json:"message"`
}

func main() {
	// Define an HTTP handler function
	http.HandleFunc("/getmessage", getMessageHandler)

	// Start the web server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// getMessageHandler is the HTTP handler function for the "/getmessage" endpoint
func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Create a Message struct
	message := Message{
		Text: "Hello, this is a JSON message!",
	}

	// Convert the Message struct to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	w.Write(jsonData)
}
