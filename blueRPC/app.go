package bluerpc

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

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

	if blueConfig.RuntimeEnv == "" {
		blueConfig.RuntimeEnv = DEVELOPMENT
	}

	fmt.Println("blueConfig.RuntimeEnv", blueConfig.RuntimeEnv)
	fiberApp := fiber.New(*fiberConfig)
	if blueConfig.RuntimeEnv == DEVELOPMENT {
		fiberApp.Use(logger.New())
	}

	return &App{
		fiberApp: fiberApp,
		config:   blueConfig,
	}
}

func NewFromApp(app *fiber.App, blueConfig Config) *App {
	var (
		tsOutputPath = "./output.ts"
		runEnv       = DEVELOPMENT
	)

	if blueConfig.OutputPath == "" {
		blueConfig.OutputPath = tsOutputPath
	}
	if blueConfig.RuntimeEnv == "" {
		blueConfig.RuntimeEnv = runEnv
	}

	return &App{
		fiberApp: app,
		config:   &blueConfig,
	}
}

func (a *App) Group(path string) *Group {

	newFiberRouter := a.fiberApp.Group(path)

	// newFiberRouter.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	return &Group{
		fiberRouter: newFiberRouter,
		basePath:    path,
		fiberApp:    a.fiberApp,
	}
}

func (a *App) Listen(port string) *App {

	// if a.config.RuntimeEnv == DEVELOPMENT {
	// 	GenerateTypes(root, a.config.OutputPath)
	// }

	a.fiberApp.Listen(port)
	return a
}
