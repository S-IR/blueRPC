package bluerpc

import (
	"github.com/gofiber/fiber/v2"
)

type Group struct {
	fiberRouter fiber.Router
	fiberApp    *fiber.App
	basePath    string
}
type ProcedureInterface interface {
	Query(h interface{}) ProcedureInterface
	Mutate(h interface{}) ProcedureInterface
}

func (grp *Group) Group(path string) *Group {
	newPath := combinePaths(grp.basePath, path)
	newFiberRouter := grp.fiberApp.Group(newPath)

	return &Group{
		fiberRouter: newFiberRouter,
		basePath:    newPath,
		fiberApp:    grp.fiberApp,
	}
}

func (grp *Group) getFiberRouter() fiber.Router {
	return grp.fiberRouter
}

func (grp *Group) Get(path string, handlers ...fiber.Handler) *Group {
	grp.fiberApp.Get(path, handlers...)
	return grp
}
