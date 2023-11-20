package bluerpc

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	//  The path where you would like the generated Typescript to be placed
	// Default is ./output.ts
	OutputPath string
}

type App struct {
	fiberApp     *fiber.App
	tsOutputPath string
}

func init() {
	root = &routeNode{
		Children: make(map[string]*routeNode),
	}
}

func New(fiberConfig *fiber.Config, blueConfig *Config) *App {
	var (
		tsOutputPath = "./output.ts"
	)

	if blueConfig.OutputPath != "" {
		tsOutputPath = blueConfig.OutputPath
	}

	return &App{
		fiberApp:     fiber.New(*fiberConfig),
		tsOutputPath: tsOutputPath,
	}
}

func NewFromApp(app *fiber.App, blueConfig Config) *App {
	var (
		tsOutputPath = "./output.ts"
	)
	if blueConfig.OutputPath != "" {
		tsOutputPath = blueConfig.OutputPath
	}

	return &App{
		fiberApp:     app,
		tsOutputPath: tsOutputPath,
	}
}

func (a *App) Group(groupPath string, procedures ...*ProcedureList) *Group {

	var plist *ProcedureList
	if len(procedures) > 0 {
		plist = procedures[0]
	}

	newFiberRouter := a.fiberApp.Group(groupPath)

	for path, proc := range *plist {
		fullPath := combinePaths(groupPath, path)
		addProcedure(a.fiberApp, fullPath, proc)
	}

	return &Group{
		fiberRouter: newFiberRouter,
		basePath:    groupPath,
		fiberApp:    a.fiberApp,
		Procedures:  *plist,
	}
}

func addProcedure[T fiber.Router](handler T, path string, proc *Procedure) {
	inputValidationMiddleware := func(c *fiber.Ctx) error {
		if proc.inputSchema == nil {
			return c.Next()
		}
		schemaType := reflect.TypeOf(proc.inputSchema)
		if schemaType.Kind() == reflect.Ptr {
			schemaType = schemaType.Elem()
		}
		schemaInstance := reflect.New(schemaType).Interface()

		// Parse the request body into the schema instance
		if err := c.BodyParser(schemaInstance); err != nil {
			fmt.Println("err here at bodyParser")
			return err
		}

		// Validate the struct
		if err := proc.validator.Struct(schemaInstance); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "input validation error",
				"details": err.Error(),
			})
		}

		return c.Next()
	}

	responseHandler := func(c *fiber.Ctx) error {
		res, err := proc.handler(c)
		if err != nil {
			fmt.Println("err here at responseHandler", err)
			return err
		}
		if proc.outputSchema != nil {
			err = validateOutput(res, proc.validator, proc.outputSchema)
			if err != nil {
				return err
			}
		}
		c.JSON(res.Body)
		return nil
	}
	AddProcedureToTree(path, proc)

	switch proc.method {
	case QUERY:
		handler.Get(path, inputValidationMiddleware, responseHandler)
	case MUTATION:
		handler.Post(path, inputValidationMiddleware, responseHandler)
	}

}

func (a *App) Listen(port string) *App {
	GenerateTypes(root, a.tsOutputPath)

	a.fiberApp.Listen(port)
	return a
}
