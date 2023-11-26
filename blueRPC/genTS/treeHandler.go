package genTypescript

import (
	"fmt"
	"strings"
)

// AddProcedureToTree adds a new procedure to the tree at the specified route.
func AddProcedureToTree(
	route string,
	queryParams, input, output interface{},
	procedureType Method) {

	segments := strings.Split(route, "/")
	currentNode := root

	for _, segment := range segments {
		if segment == "" {
			continue
		}
		if currentNode.Children == nil {
			currentNode.Children = map[string]*routeNode{}
		}

		if _, exists := currentNode.Children[segment]; !exists {
			currentNode.Children[segment] = &routeNode{
				Children: nil,
			}
		}
		currentNode = currentNode.Children[segment]

	}

	if procedureType == QUERY {
		currentNode.query = &procedureInfo{
			queryParams: queryParams,
			input:       nil,
			output:      output,
		}
	} else if procedureType == MUTATION {
		currentNode.mutation = &procedureInfo{
			queryParams: queryParams,
			input:       input,
			output:      output,
		}
	}
}
func moveProcedure(oldRoute string, newRoute string) {
	oldRouteSegments := strings.Split(oldRoute, "/")
	oldNode := root
	for _, segment := range oldRouteSegments {
		if oldNode.Children == nil || oldNode.Children[segment] == nil {
			fmt.Println("This procedure route does not exist")
			return
		}
		oldNode = oldNode.Children[segment]

	}

	newRouteSegments := strings.Split(newRoute, "/")
	newNode := root

	for _, segment := range newRouteSegments {
		if segment == "" {
			continue
		}
		if newNode.Children == nil {
			newNode.Children = map[string]*routeNode{}
		}

		if _, exists := newNode.Children[segment]; !exists {
			newNode.Children[segment] = &routeNode{
				Children: nil,
			}
		}
		newNode = newNode.Children[segment]
	}

	newNode = oldNode
	oldNode = nil
}
