package pastis

import (
	"strings"
)

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handler
}

func (n *node) addNode(method, path string, handler Handler) {
	components := strings.Split(strings.Trim(path, componentsSeparator), componentsSeparator)
	count := len(components)
	for {
		aNode, component := n.traverse(components, nil)
		if aNode.component == component && count == 1 {
			aNode.methods[method] = handler
			return
		}
		newNode := node{
			component:    component,
			isNamedParam: false,
			methods:      make(map[string]Handler),
		}
		if len(component) > 0 && string(component[0]) == pathParamIndicator {
			newNode.isNamedParam = true
		}
		if count == 1 {
			newNode.methods[method] = handler
		}
		aNode.children = append(aNode.children, &newNode)
		count--
		if count == 0 {
			break
		}
	}
}

func (n *node) traverse(components []string, params params) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.add(strings.Trim(child.component, pathParamIndicator), component)
				}
				next := components[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}