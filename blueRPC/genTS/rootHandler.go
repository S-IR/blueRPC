package genTypescript

import (
	"fmt"
	"os"
	"strings"
)

type Method string

var (
	QUERY    Method = "query"
	MUTATION Method = "mutation"
)

// procedureInfo is a generic struct that embeds procedureInfoBase.
// MAYBE this should be different for QUERY and MUTATION since QUERY has no INPUT
type procedureInfo struct {
	queryParams interface{}
	input       interface{}
	output      interface{}
}

// routeNode is a struct to represent a node in the tree.
type routeNode struct {
	Children map[string]*routeNode
	query    *procedureInfo
	mutation *procedureInfo
}

// root is the root node of the tree.
var root *routeNode

func init() {
	root = &routeNode{
		Children: make(map[string]*routeNode),
		query:    nil,
		mutation: nil,
	}
}

func StartGenerating(filePath, name, startingPath string) error {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("type %s =", name))
	startingPath = strings.ReplaceAll(startingPath, "/", "")

	if startingPath == "" || startingPath == "/" {
		nodeToTS(&builder, root, true)
	} else if root.Children != nil && root.Children[startingPath] != nil {
		nodeToTS(&builder, root.Children[startingPath], true)
	} else {
		fmt.Println("root.Children", root.Children)

		for path, _ := range root.Children {
			fmt.Println("PATH OF ROUTE " + path)
		}
		panic("This starting path does not exist or is not set" + startingPath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(builder.String())
	return err

}
