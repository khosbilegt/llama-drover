package handlers

import (
	"io"
	"log"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func HandleGeneratePrompt(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Print the body content
	log.Printf("Request body: %s", string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Prompt generated successfully"))
}
