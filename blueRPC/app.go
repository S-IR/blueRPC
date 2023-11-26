package bluerpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"strings"
	"time"

	genTypescript "github.com/S-IR/blueRPC/blueRPC/genTS"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp"
)

type validatorFn func(interface{}) error
type ValidRouter interface {
	getFiberRouter() *fiber.Router
	getPath() string
}

type App struct {
	fiberApp   *fiber.App
	config     *Config
	startGroup *fiber.Router
}

func (a *App) getFiberRouter() *fiber.Router {
	return a.startGroup
}

func New(blueConfig ...*Config) *App {

	var cfg *Config

	if len(blueConfig) > 0 {
		cfg = blueConfig[0]
	} else {
		cfg = &Config{}
	}

	if cfg.FiberConfig == nil {
		cfg.FiberConfig = &fiber.Config{}
	}

	fiberApp := fiber.New(*cfg.FiberConfig)
	cfg, childApp, startGroup := setAppDefaults(blueConfig, fiberApp)

	return &App{
		fiberApp:   childApp,
		config:     cfg,
		startGroup: startGroup,
	}
}

func NewFromApp(fiberApp *fiber.App, blueConfig ...*Config) *App {
	var cfg *Config

	if len(blueConfig) > 0 {
		cfg = blueConfig[0]
	} else {
		cfg = &Config{}
	}
	cfg, childApp, startGroup := setAppDefaults(blueConfig, fiberApp)

	return &App{
		fiberApp:   childApp,
		config:     cfg,
		startGroup: startGroup,
	}
}

func setAppDefaults(blueConfig []*Config, fiberApp *fiber.App) (*Config, *fiber.App, *fiber.Router) {

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
	startGroup := fiberApp.Group(cfg.StartingPath)
	if !cfg.DisableRequestLogging {
		fiberApp.Use(logger.New())
	}
	if !cfg.DisableJSONOnlyErrors {
		fiberApp.Use(DefaultErrorMiddleware)
	}

	return cfg, fiberApp, &startGroup
}

func (a *App) Group(path string) *Group {

	newFiberRouter := (*a.startGroup).Group(path)

	// newFiberRouter.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	return &Group{
		fiberRouter: newFiberRouter,
		basePath:    a.config.StartingPath + path,
		fiberApp:    a.fiberApp,
	}
}

func (a *App) Static(route, filePath string, settings ...*fiber.Static) *App {

	var actualSettings *fiber.Static
	if len(settings) > 1 {
		actualSettings = settings[0]
	} else {
		actualSettings = nil
	}
	a.fiberApp.Static(route, filePath, *actualSettings)
	return a
}
func (a *App) Server() *fasthttp.Server {
	return a.fiberApp.Server()
}
func (a *App) Shutdown() error {
	return a.fiberApp.Shutdown()
}
func (a *App) ShutdownWithTimeout(timeout time.Duration) error {
	return a.fiberApp.ShutdownWithTimeout(timeout)
}

func (a *App) ShutdownWithContext(ctx context.Context) error {
	return a.fiberApp.ShutdownWithContext(ctx)
}
func (a *App) HandlersCount() uint32 {
	return a.fiberApp.HandlersCount()
}
func (a *App) Stack() [][]*fiber.Route {
	return a.fiberApp.Stack()
}
func (a *App) Name(name string) ValidRouter {
	a.fiberApp.Name(name)
	return a
}
func (a *App) GetRoutes(filterUseOption ...bool) []fiber.Route {
	return a.fiberApp.GetRoutes(filterUseOption...)
}

// same as the Config method in fiber
func (a *App) FiberConfig() fiber.Config {
	return a.fiberApp.Config()
}
func (a *App) BluerpcConfig() Config {
	return *a.config
}

// same as the Listen method in fiber
func (a *App) Listen(port string) *App {

	var name string

	lastSlashIndex := strings.LastIndex(a.config.StartingPath, "/")
	if lastSlashIndex == -1 {
		name = a.config.StartingPath
	} else {
		name = a.config.StartingPath[lastSlashIndex+1:]
	}

	if a.config.disableGenerateTS == false {
		// start := time.Now()
		err := genTypescript.StartGenerating(a.config.OutputPath, name, a.config.StartingPath)
		if err != nil {
			panic(err)
		}
		// elapsed := time.Since(start)
		// fmt.Printf(fiber.DefaultColors.Green+"Execution time for GENERATING TYPESCRIPT: %s\n"+fiber.DefaultColors.Reset, elapsed)
	}

	// routes := a.fiberApp.GetRoutes()
	// fmt.Println("Registered Routes:")

	// for _, route := range routes {
	// 	fmt.Printf("%s %s\n", route.Method, route.Path)
	// }

	a.fiberApp.Listen(port)
	return a
}

// sale as ListenTLS method in fiber
func (a *App) ListenTLS(addr, certFile, keyFile string) error {
	return a.ListenTLS(addr, certFile, keyFile)
}

// sale as Listener method in fiber
func (a *App) Listener(ln net.Listener) error {
	return a.fiberApp.Listener(ln)
}

// sale as Test method in fiber
func (a *App) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return a.fiberApp.Test(req, msTimeout...)
}
func (a *App) Hooks() *fiber.Hooks {
	return a.fiberApp.Hooks()
}

// sale as ListenMutualTLSWithCertificate method in fiber
func (a *App) ListenMutualTLSWithCertificate(addr string, cert tls.Certificate, clientCertPool *x509.CertPool) error {
	return a.fiberApp.ListenMutualTLSWithCertificate(addr, cert, clientCertPool)
}
func (a *App) getPath() string {
	return a.config.StartingPath
}
