package bluerpc

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
func isLoggerMiddleware(middleware func(*fiber.Ctx) error) bool {
	// Get the name of the middleware function
	middlewareName := getFunctionName(middleware)
	// Get the name of the logger middleware function
	loggerName := getFunctionName(logger.New())
	// Compare the names
	return middlewareName == loggerName
}

// Function to get the name of a function
func getFunctionName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	// Extract only the function name
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}
