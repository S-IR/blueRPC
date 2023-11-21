package bluerpc

// func GenerateTypes(node *routeNode, outputPath string) {

// 	fmt.Println("node", node)
// 	panic("wait here")
// 	var tsCode strings.Builder
// 	generateTypesRecursive(node, "", &tsCode)

// 	// Create or open the TypeScript file
// 	file, err := os.Create(outputPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// Write the generated TypeScript code to the file
// 	_, err = file.WriteString(tsCode.String())
// 	if err != nil {
// 		panic(err)
// 	}
// }
// func generateTypesRecursive(node *routeNode, path string, tsCode *strings.Builder) {
// 	// Iterate over children
// 	for segment, child := range node.Children {
// 		newPath := path
// 		if path != "" {
// 			newPath += "."
// 		}
// 		newPath += segment

// 		// Recursively generate types for children
// 		generateTypesRecursive(child, newPath, tsCode)
// 	}

// 	// If the current node has a procedure, generate TypeScript code for it
// 	if node.Procedure != nil {
// 		inputTS := generateTSFields(node.Procedure.inputSchema)
// 		outputTS := generateTSFields(node.Procedure.outputSchema)

// 		// Generate TypeScript function based on method
// 		functionName := "query"
// 		if node.Procedure.method == MUTATION {
// 			functionName = "mutation"
// 		}

// 		// Construct the TypeScript function type
// 		tsFunction := fmt.Sprintf("const %s: (input: {%s}) => {%s};", functionName, inputTS, outputTS)
// 		tsCode.WriteString(tsFunction + "\n")
// 	}
// }
// func generateTSFields(schema interface{}) string {
// 	var fields []string
// 	t := reflect.TypeOf(schema)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		tsType := "string"
// 		fields = append(fields, fmt.Sprintf("%s: %s", field.Name, tsType))
// 	}

// 	return strings.Join(fields, ", ")
// }
