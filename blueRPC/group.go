package bluerpc

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type Group struct {
	fiberRouter fiber.Router
	fiberApp    *fiber.App
	basePath    string
	Procedures  map[string]*Procedure
	groupInput  interface{}
	groupOutput interface{}
}

type ProcedureList map[string]*Procedure

func (grp *Group) Group(path string, procedures *ProcedureList) *Group {
	newPath := combinePaths(grp.basePath, path)
	newFiberRouter := grp.fiberApp.Group(newPath)

	for path, proc := range *procedures {
		addProcedure(grp.fiberRouter, path, proc)
	}

	return &Group{
		fiberRouter: newFiberRouter,
		basePath:    newPath,
		fiberApp:    grp.fiberApp,
		groupInput:  grp.groupInput,
		groupOutput: grp.groupOutput,
		Procedures:  *procedures,
	}
}

func (grp *Group) Get(path string, handlers ...fiber.Handler) *Group {
	grp.fiberApp.Get(path, handlers...)
	return grp
}

func addProcedure[T fiber.Router](handler T, path string, proc *Procedure) {

	validatorFn := *proc.validatorFn

	inputValidationMiddleware := func(c *fiber.Ctx) error {
		if proc.inputSchema == nil || proc.validatorFn == nil {
			return c.Next()
		}

		localSchema := copySchema(proc.inputSchema)

		// Parse the request body into the schema instance
		if err := c.BodyParser(localSchema); err != nil {
			fmt.Println("err here at bodyParser")
			return err
		}

		// Validate the struct
		fmt.Println("arrived here at validate the struct", localSchema)
		if err := validatorFn(localSchema); err != nil {
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
		if proc.outputSchema != nil && proc.validatorFn != nil {
			err = validateOutput(res.Body, validatorFn, proc.outputSchema)
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

func validateOutput(resBody interface{}, validatorFn validatorFn, outputSchema interface{}) error {

	if reflect.TypeOf(resBody) != reflect.TypeOf(outputSchema) {
		log.Fatalf("Response body and the provided output schema are of different types. Response body is %s while Output Schema is %s", reflect.TypeOf(resBody), reflect.TypeOf(outputSchema))
	}

	if err := validatorFn(resBody); err != nil {
		panic(err)
	}
	return nil
}
