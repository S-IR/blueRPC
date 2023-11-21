package bluerpc

import (
	"reflect"
	"strings"
)

// combinePaths takes two route strings and combines them into one.
func combinePaths(route1, route2 string) string {
	// Ensure both routes start and end without a slash.
	cleanRoute1 := strings.TrimSuffix(route1, "/")
	cleanRoute2 := strings.TrimPrefix(route2, "/")

	// Combine the routes with a single slash in between.
	fullRoute := cleanRoute1 + "/" + cleanRoute2
	return fullRoute
}

func copySchema(schema interface{}) any {
	schemaType := reflect.TypeOf(schema)
	if schemaType.Kind() == reflect.Ptr {
		schemaType = schemaType.Elem()
	}
	localSchema := reflect.New(schemaType).Interface()
	return localSchema
}
