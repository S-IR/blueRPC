package genTypescript

import (
	"fmt"
	"sort"
	"strings"
)

func nodeToTS(stringBuilder *strings.Builder, node *routeNode, isLast bool) {
	stringBuilder.WriteString("{")

	if node.query != nil {
		stringBuilder.WriteString(fmt.Sprintf("query:"))
		queryParams, output := node.query.queryParams, node.query.input
		genTSFuncFromQuery(stringBuilder, queryParams, output)

	}

	if node.mutation != nil {
		stringBuilder.WriteString(fmt.Sprintf("mutation:"))
		queryParams, input, output := node.mutation.queryParams, node.mutation.input, node.mutation.output
		genTSFuncFromMutation(stringBuilder, queryParams, input, output)

	}

	if node.Children != nil {

		keys := getSortedKeys(node.Children)
		for i, path := range keys {
			stringBuilder.WriteString(fmt.Sprintf("%s:", path))
			nodeToTS(stringBuilder, node.Children[path], i == len(keys)-1)
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
