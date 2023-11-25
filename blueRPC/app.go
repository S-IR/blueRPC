package bluerpc

import (
	"fmt"
	"strings"
	"time"

	genTypescript "github.com/S-IR/blueRPC/blueRPC/genTS"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type validatorFn func(interface{}) error
type ValidRouter interface {
	getFiberRouter() fiber.Router
	getPath() string
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
func (a *App) Listen(port string) *App {

	var name string
	lastSlashIndex := strings.LastIndex(a.config.StartingPath, "/")
	if lastSlashIndex == -1 {
		name = a.config.StartingPath
	} else {
		name = a.config.StartingPath[lastSlashIndex+1:]
	}

	if a.config.disableGenerateTS == false {
		start := time.Now()
		err := genTypescript.StartGenerating(a.config.OutputPath, name)
		if err != nil {
			panic(err)
		}
		elapsed := time.Since(start)
		fmt.Printf(fiber.DefaultColors.Green+"Execution time for GENERATING TYPESCRIPT: %s\n", elapsed)
	}

	a.fiberApp.Listen(port)
	return a
}
func (a *App) getPath() string {
	return a.config.StartingPath
}
