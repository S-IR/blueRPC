package genTypescript

import (
	"fmt"
	"sort"
	"strings"
)

func nodeToTS(stringBuilder *strings.Builder, node *routeNode, isLast bool, currentPath string) {

	stringBuilder.WriteString("{")

	if node.query != nil {
		stringBuilder.WriteString(fmt.Sprintf("query: async "))
		queryParams, output := node.query.queryParams, node.query.output
		genTSFuncFromQuery(stringBuilder, queryParams, output, currentPath)

	}

	if node.mutation != nil {
		stringBuilder.WriteString(fmt.Sprintf("mutation: async "))
		queryParams, input, output := node.mutation.queryParams, node.mutation.input, node.mutation.output
		genTSFuncFromMutation(stringBuilder, queryParams, input, output, currentPath)

	}

	if node.Children != nil {

		keys := getSortedKeys(node.Children)
		for i, path := range keys {
			stringBuilder.WriteString(fmt.Sprintf("%s:", path))
			nodeToTS(stringBuilder, node.Children[path], i == len(keys)-1, currentPath+"/"+path)
		}

	}

	stringBuilder.WriteString("}")
	if !isLast {
		stringBuilder.WriteString(",")
	}
}
func getSortedKeys(m map[string]*routeNode) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
