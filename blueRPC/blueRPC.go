package bluerpc

type Config struct {
	//  The path where you would like the generated Typescript to be placed
	// Default is ./output.ts
	OutputPath string

	ValidatorFn validatorFn
}

func init() {
	root = &routeNode{
		Children: make(map[string]*routeNode),
	}
}
