package tree

func (n Node) Intersection(node Node) Node {
	if n.IsFolder() && node.IsFolder() {
		if n.Name != node.Name {
			return NewEmptyNode()
		}

		var children []Node

		for i, child := range n.Children {
			_, isPresent := node.Children[i]

			if !isPresent {
				continue
			}

			intersection := child.Intersection(node.Children[i])
			children = append(children, intersection)
		}
		return NewFolderNode(n.Name, children)
	}

	if n.IsFile && node.IsFile {
		if n.Name == node.Name {
			return n
		}
		return NewEmptyNode()
	}

	return NewEmptyNode()
}
