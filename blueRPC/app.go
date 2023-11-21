package bluerpc

import "github.com/gofiber/fiber/v2"

type validatorFn func(interface{}) error

type App struct {
	fiberApp *fiber.App

	config *Config
}

func New(fiberConfig *fiber.Config, blueConfig *Config) *App {
	var (
		tsOutputPath = "./output.ts"
	)

	if blueConfig.OutputPath == "" {
		blueConfig.OutputPath = tsOutputPath
	}

	return &App{
		fiberApp: fiber.New(*fiberConfig),
		config:   blueConfig,
	}
}

func NewFromApp(app *fiber.App, blueConfig Config) *App {
	var (
		tsOutputPath = "./output.ts"
	)

	if blueConfig.OutputPath == "" {
		blueConfig.OutputPath = tsOutputPath
	}

	return &App{
		fiberApp: app,
		config:   &blueConfig,
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

func (a *App) Listen(port string) *App {
	// GenerateTypes(root, a.tsOutputPath)

	a.fiberApp.Listen(port)
	return a
}
