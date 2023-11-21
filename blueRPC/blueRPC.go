package bluerpc

type RunEnv string

const (
	DEVELOPMENT RunEnv = "development"
	PRODUCTION  RunEnv = "production"
)

type Config struct {
	//  The path where you would like the generated Typescript to be placed
	// Default is ./output.ts
	OutputPath string

	//The Runtime environment you are in. Either development or production
	//
	RuntimeEnv RunEnv

	ValidatorFn validatorFn
}

func init() {
	root = &routeNode{
		Children: make(map[string]*routeNode),
	}
}
