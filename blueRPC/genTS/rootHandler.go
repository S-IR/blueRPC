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

func StartGenerating(filePath, name string) error {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("type %s =", name))

	nodeToTS(&builder, root, true)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(builder.String())
	return err

}
