package bluerpc

import (
	"fmt"

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

func (grp *Group) Get(path string, handlers ...fiber.Handler) *Group {
	grp.fiberApp.Get(path, handlers...)
	return grp
}

func addProcedure[T fiber.Router, input any, output any](handler T, path string, proc *Procedure[input, output]) {

	validatorFn := *proc.validatorFn
	validateInput := func(c *fiber.Ctx) (input, error) {

		inputInstance := new(input)
		if proc.inputSchema == nil {
			return *inputInstance, nil
		}

		if err := c.BodyParser(inputInstance); err != nil {
			fmt.Println("err here at bodyParser")
			return *inputInstance, err
		}
		// Validate the struct
		fmt.Println("arrived here at validate the struct", inputInstance)
		if err := validatorFn(inputInstance); err != nil {

			return *inputInstance, &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: err.Error(),
			}

		}
		return *inputInstance, nil
	}

	validateOutput := func(res Res[output]) error {
		if proc.outputSchema == nil {
			return nil
		}

		if err := validatorFn(res.Body); err != nil {
			panic(err.Error())
		}
		return nil
	}

	FullHandler := func(c *fiber.Ctx) error {
		input, err := validateInput(c)
		if err != nil {
			return err
		}
		res, err := proc.handler(c, input)

		if err != nil {
			return err
		}
		err = validateOutput(res)

		if err != nil {
			return err
		}

		return c.JSON(res.Body)
	}

	procInfo := procedureInfo{
		path:   path,
		input:  proc.inputSchema,
		output: proc.outputSchema,
	}
	AddProcedureToTree(procInfo)

	switch proc.method {
	case QUERY:
		handler.Get(path, FullHandler)
	case MUTATION:
		handler.Post(path, FullHandler)
	default:
		panic(fmt.Sprintf("This Procedure has no method, at %s", path))
	}

}
