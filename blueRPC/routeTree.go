package bluerpc

import (
	"fmt"
	"strings"
)

type routeNode struct {
	Children  map[string]*routeNode
	Procedure *Procedure
}

var (
	root *routeNode
)

func AddProcedureToTree(fullPath string, proc *Procedure) {
	segments := strings.Split(fullPath, "/")

	fmt.Println("root", root)
	currentNode := root
	for _, segment := range segments {
		if segment == "" {
			continue
		}

		if _, exists := currentNode.Children[segment]; !exists {
			currentNode.Children[segment] = &routeNode{
				Children: make(map[string]*routeNode),
			}
		}
		currentNode = currentNode.Children[segment]
	}

	currentNode.Procedure = proc
}
