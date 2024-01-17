package tree

func (n Node) Difference(node Node) Node {
	if n.IsFolder() && node.IsFolder() {
		var children []Node
		for i, child := range n.Children {
			_, isPresent := node.Children[i]

			if !isPresent {
				children = append(children, child)
				continue
			}

			difference := child.Difference(node.Children[i])

			if len(difference.Children) != 0 {
				children = append(children, difference)
			}
		}
		return NewFolderNode(n.Name, children)
	}

	if n.IsFile && node.IsFile {
		if n.IsEquals(node) {
			return NewEmptyNode()
		}
		return n
	}

	return n
}
