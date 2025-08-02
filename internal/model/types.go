package model

type Cluster struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Model string `json:"model" bson:"model"`
	Nodes []Node `json:"nodes" bson:"nodes"`
}

type Node struct {
	ID        string `json:"id" bson:"id"`
	ClusterID string `json:"cluster_id" bson:"cluster_id"`
	Name      string `json:"name" bson:"name"`
	Status    string `json:"status" bson:"status"`
	IPAddress string `json:"ip_address" bson:"ip_address"`
	Port      int    `json:"port" bson:"port"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
