package bluerpc

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type validatorFn func(interface{}) error
type ValidRouter interface {
	getFiberRouter() fiber.Router
}

type App struct {
	fiberApp *fiber.App
	config   *Config
}

func (a *App) getFiberRouter() fiber.Router {
	return a.fiberApp
}

func New(blueConfig ...*Config) *App {

	cfg := setAppDefaults(blueConfig)

	fiberApp := fiber.New(*cfg.FiberConfig)
	if !cfg.DisableRequestLogging {
		fiberApp.Use(logger.New())
	}

	return &App{
		fiberApp: fiberApp,
		config:   cfg,
	}
}

func NewFromApp(app *fiber.App, blueConfig ...*Config) *App {

	cfg := setAppDefaults(blueConfig)

	return &App{
		fiberApp: app,
		config:   cfg,
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
func setAppDefaults(blueConfig []*Config) *Config {
	var cfg *Config

	if len(blueConfig) > 0 {
		cfg = blueConfig[0]
	} else {
		cfg = &Config{}
	}
	if cfg.FiberConfig == nil {
		cfg.FiberConfig = &fiber.Config{}
	}

	var (
		tsOutputPath = "./output.ts"
		startPath    = "/bluerpc"
	)

	if cfg.OutputPath == "" {
		cfg.OutputPath = tsOutputPath
	}

	if cfg.StartingPath == "" {
		cfg.StartingPath = startPath
	}
	return cfg
}
