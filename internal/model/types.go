package model

type Herd struct {
	Name  string
	Model string
	Nodes []Node
}

type Node struct {
	Name      string
	Status    string
	IPAddress string
	Port      int
}
