package tree

type Node struct {
	Name     string
	Children map[string]Node
	IsFile   bool
}

func NewEmptyNode() Node {
	return Node{
		Name:     "",
		Children: map[string]Node{},
		IsFile:   false,
	}
}

func NewFileNode(name string) Node {
	return Node{
		Name:   name,
		IsFile: true,
	}
}

func NewFolderNode(name string, children []Node) Node {
	childrenAsMap := map[string]Node{}

	for _, child := range children {
		childrenAsMap[child.Name] = child
	}

	return Node{
		Name:     name,
		Children: childrenAsMap,
		IsFile:   false,
	}
}

func (n Node) IsFolder() bool {
	return !n.IsFile
}

func (n Node) Print() {
	n.print("")
}

func (n Node) print(space string) {
	if n.IsFile {
		println(space + "File: " + n.Name)
		return
	}
	println(space + "Folder: " + n.Name)
	for _, child := range n.Children {
		child.print(space + "	")
	}

}

func (n Node) IsEquals(node Node) bool {
	if n.Name != node.Name {
		return false
	}
	if n.IsFile != node.IsFile {
		return false
	}
	if len(n.Children) != len(node.Children) {
		return false
	}
	for i, child := range n.Children {
		if !child.IsEquals(node.Children[i]) {
			return false
		}
	}
	return true
}

func (n Node) Count() int {
	if n.IsFile {
		return 1
	}
	count := 0
	for _, child := range n.Children {
		count += child.Count()
	}
	return count
}
