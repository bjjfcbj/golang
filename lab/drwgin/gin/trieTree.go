package gin

import (
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	iswild   bool
}

func (n *node) matchchild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.iswild {
			return child
		}
	}
	return nil
}

func (n *node) matchchildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.iswild {
			nodes=append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchchild(part)
	if child == nil {
		child = &node{
			part:   part,
			iswild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchchildren(part)

	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}
