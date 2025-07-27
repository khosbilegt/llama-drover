package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/khosbilegt/llama-drover/internal/coordinator"
	"github.com/khosbilegt/llama-drover/internal/model"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResp := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}

func writeJSONSuccess(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	successResp := SuccessResponse{
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(successResp)
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSONSuccess(w, "Service is healthy", nil, http.StatusOK)
}

func HandleGeneratePrompt(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		writeJSONError(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Request body: %s", string(body))

	writeJSONSuccess(w, "Prompt generated successfully", nil, http.StatusOK)
}

func HandleCreateHerd(w http.ResponseWriter, r *http.Request) {
	var herd model.Herd
	err := json.NewDecoder(r.Body).Decode(&herd)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if herd.Name == "" || herd.Model == "" {
		log.Println("Herd name or model is empty")
		writeJSONError(w, "Herd name and model must be provided", http.StatusBadRequest)
		return
	}

	herd, createErr := coordinator.CreateHerd(herd.Name, herd.Model)
	if createErr != nil {
		log.Printf("Error creating herd: %v", createErr)
		writeJSONError(w, "Error creating herd", http.StatusInternalServerError)
		return
	}

	writeJSONSuccess(w, "Herd created successfully", herd, http.StatusCreated)
}

func HandleDeleteHerd(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/herd/"), "/")
	log.Println("Fetching herd details", parts)
	var herdID string = parts[0]
	if herdID == "" {
		writeJSONError(w, "Herd ID must be provided", http.StatusBadRequest)
		return
	}

	coordinator.DeleteHerd(herdID)

	writeJSONSuccess(w, "Herd deleted successfully", nil, http.StatusOK)
}

func HandleGetHerd(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/herd/"), "/")
	log.Println("Fetching herd details", parts)
	var herdID string = parts[0]
	if herdID == "" {
		writeJSONError(w, "Herd ID must be provided", http.StatusBadRequest)
		return
	}

	herd, err := coordinator.GetHerd(herdID)
	if err != nil {
		log.Printf("Error getting herd: %v", err)
		writeJSONError(w, "Error getting herd", http.StatusInternalServerError)
		return
	}

	writeJSONSuccess(w, "Herd details fetched successfully", herd, http.StatusOK)
}

func HandleListHerds(w http.ResponseWriter, r *http.Request) {
	herds, err := coordinator.ListHerds()
	if err != nil {
		log.Printf("Error listing herds: %v", err)
		writeJSONError(w, "Error listing herds", http.StatusInternalServerError)
		return
	}
	herdsNormalized := herds
	if (herdsNormalized) == nil {
		herdsNormalized = []model.Herd{}
	}
	writeJSONSuccess(w, "Herds fetched successfully", herdsNormalized, http.StatusOK)
}

func HandleFetchNode(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/node/"), "/")
	nodeId := parts[0]
	if nodeId == "" {
		writeJSONError(w, "Node ID must be provided", http.StatusBadRequest)
		return
	}
	node, err := coordinator.GetNode(nodeId)
	if err != nil {
		log.Printf("Error fetching node: %v", err)
		writeJSONError(w, "Error fetching node", http.StatusInternalServerError)
		return
	}

	writeJSONSuccess(w, "Node details fetched successfully", node, http.StatusOK)
}

func HandleListNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := coordinator.ListNodes()
	if err != nil {
		log.Printf("Error listing nodes: %v", err)
		writeJSONError(w, "Error listing nodes", http.StatusInternalServerError)
		return
	}

	nodesNormalized := nodes
	if nodesNormalized == nil {
		nodesNormalized = []model.Node{}
	}
	writeJSONSuccess(w, "Nodes fetched successfully", nodesNormalized, http.StatusOK)
}

func HandleRegisterNode(w http.ResponseWriter, r *http.Request) {
	var nodeInput model.Node
	err := json.NewDecoder(r.Body).Decode(&nodeInput)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if nodeInput.Name == "" {
		log.Println("Node name or ID is empty")
		writeJSONError(w, "Node name must be provided", http.StatusBadRequest)
		return
	}
	node, err := coordinator.CreateNode(nodeInput)
	if err != nil {
		log.Printf("Error registering node: %v", err)
		writeJSONError(w, "Error registering node", http.StatusInternalServerError)
		return
	}
	writeJSONSuccess(w, "Node registered successfully", node, http.StatusCreated)
}

func HandleDeleteNode(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/node/"), "/")
	nodeId := parts[0]
	if nodeId == "" {
		writeJSONError(w, "Node ID must be provided", http.StatusBadRequest)
		return
	}
	err := coordinator.DeleteNode(nodeId)
	if err != nil {
		log.Printf("Error deleting node: %v", err)
		writeJSONError(w, "Error deleting node", http.StatusInternalServerError)
		return
	}

	writeJSONSuccess(w, "Node deleted successfully", nil, http.StatusOK)
}
