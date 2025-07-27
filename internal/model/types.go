package model

type Herd struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Nodes []Node `json:"nodes"`
}

type Node struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	IPAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
