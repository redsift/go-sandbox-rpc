package sift

type Root struct {
	Dag *Dag `json:"dag"`
}

type Dag struct {
	Nodes []Node `json:"nodes"`
}

func (s *Root) HasDag() bool {
	if s.Dag != nil && len(s.Dag.Nodes) > 0 {
		return true
	}

	return false
}

// Node related
// -----

// Paths relative to the directory containing the sift.json
type Implementation struct {
	Go string `json:"go,omitempty"` // path to .go implementation
}

type Node struct {
	Description    string          `json:"#,omitempty"`
	Implementation *Implementation `json:"implementation,omitempty"`
}

func (n *Node) Impl() string {
	if n.Implementation == nil || len(n.Implementation.Go) == 0 {
		return ""
	}
	return n.Implementation.Go
}
