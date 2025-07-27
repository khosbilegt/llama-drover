package model

type Herd struct {
	ID    string
	Name  string
	Model string
	Nodes []Node
}

type Node struct {
	ID        string
	Name      string
	Status    string
	IPAddress string
	Port      int
}
