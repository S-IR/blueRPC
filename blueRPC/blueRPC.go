package bluerpc

import "github.com/gofiber/fiber/v2"

type RuntimeEnv string

const (
	DEVELOPMENT RuntimeEnv = "development"
	PRODUCTION  RuntimeEnv = "production"
)

type Config struct {
	FiberConfig *fiber.Config
	//  The path where you would like the generated Typescript to be placed.
	// Default is ./output.ts
	OutputPath string

	// Boolean that determines if any typescript types will be generated.
	// Default is false. Set this to TRUE in production
	disableGenerateTS bool

	//The function that will be used to validate your struct fields.
	ValidatorFn validatorFn

	//Disables the fiber logger middleware that is added.
	//False by default. Set this to TRUE in production
	DisableRequestLogging bool

	//This is the home route that all of the bluerpc routes will start from. Default is "/bluerpc".
	StartingPath string

	// by default Bluerpc transforms every error that is sent to the user into a json by the format  DefaultErrorResponse. Turn this to true if you would like fiber to handle what form will the errors have or write your own middleware to convert all of the errors to your liking
	DisableJSONOnlyErrors bool
}
