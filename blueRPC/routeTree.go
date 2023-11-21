package bluerpc

import (
	"fmt"
	"strings"
)

type procedureInfo struct {
	path   string
	input  interface{}
	output interface{}
}

type routeNode struct {
	Children      map[string]*routeNode
	procedureInfo procedureInfo
}

var (
	root *routeNode
)

func AddProcedureToTree(procInfo procedureInfo) {
	segments := strings.Split(procInfo.path, "/")

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

	currentNode.procedureInfo = procInfo
}
