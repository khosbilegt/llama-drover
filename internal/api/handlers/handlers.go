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

func HandleCreateCluster(w http.ResponseWriter, r *http.Request) {
	var cluster model.Cluster
	err := json.NewDecoder(r.Body).Decode(&cluster)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if cluster.Name == "" || cluster.Model == "" {
		log.Println("Cluster name or model is empty")
		writeJSONError(w, "Cluster name and model must be provided", http.StatusBadRequest)
		return
	}

	cluster, createErr := coordinator.CreateCluster(cluster.Name, cluster.Model)
	if createErr != nil {
		log.Printf("Error creating Cluster: %v", createErr)
		writeJSONError(w, "Error creating Cluster", http.StatusInternalServerError)
		return
	}

	writeJSONSuccess(w, "Cluster created successfully", cluster, http.StatusCreated)
}

func HandleDeleteCluster(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/cluster/"), "/")
	log.Println("Fetching Cluster details", parts)
	var clusterID string = parts[0]
	if clusterID == "" {
		writeJSONError(w, "Cluster ID must be provided", http.StatusBadRequest)
		return
	}

	coordinator.DeleteCluster(clusterID)

	writeJSONSuccess(w, "Cluster deleted successfully", nil, http.StatusOK)
}

func HandleGetCluster(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/cluster/"), "/")
	log.Println("Fetching Cluster details", parts)
	var clusterID string = parts[0]
	if clusterID == "" {
		writeJSONError(w, "Cluster ID must be provided", http.StatusBadRequest)
		return
	}

	Cluster, err := coordinator.GetCluster(clusterID)
	if err != nil {
		log.Printf("Error getting Cluster: %v", err)
		writeJSONError(w, "Error getting Cluster", http.StatusInternalServerError)
		return
	}
	if Cluster.ID == "" {
		log.Printf("Cluster with ID %s not found", clusterID)
		writeJSONError(w, "Cluster not found", http.StatusNotFound)
		return
	}

	writeJSONSuccess(w, "Cluster details fetched successfully", Cluster, http.StatusOK)
}

func HandleListClusters(w http.ResponseWriter, r *http.Request) {
	clusters, err := coordinator.ListClusters()
	if err != nil {
		log.Printf("Error listing clusters: %v", err)
		writeJSONError(w, "Error listing clusters", http.StatusInternalServerError)
		return
	}
	clustersNormalized := clusters
	if (clustersNormalized) == nil {
		clustersNormalized = []model.Cluster{}
	}
	writeJSONSuccess(w, "clusters fetched successfully", clustersNormalized, http.StatusOK)
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
	if node.ID == "" {
		log.Printf("Node with ID %s not found", nodeId)
		writeJSONError(w, "Node not found", http.StatusNotFound)
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

func HandleCreateNode(w http.ResponseWriter, r *http.Request) {
	var nodeInput model.Node
	err := json.NewDecoder(r.Body).Decode(&nodeInput)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if nodeInput.Name == "" || nodeInput.ClusterID == "" {
		log.Println("Node name or Cluster ID is empty")
		writeJSONError(w, "Node name and Cluster ID must be provided", http.StatusBadRequest)
		return
	}

	cluster, clusterErr := coordinator.GetCluster(nodeInput.ClusterID)
	if clusterErr != nil {
		log.Printf("Error fetching cluster: %v", clusterErr)
		writeJSONError(w, "Error fetching cluster", http.StatusInternalServerError)
		return
	}
	if cluster.ID == "" {
		log.Printf("Cluster with ID %s does not exist", nodeInput.ClusterID)
		writeJSONError(w, "Cluster does not exist", http.StatusNotFound)
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
