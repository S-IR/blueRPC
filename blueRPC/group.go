package bluerpc

import (
	"encoding/json"
	"reflect"

	"github.com/go-playground/validator"
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
func validateOutput(resBody interface{}, validate *validator.Validate, outputSchema interface{}) error {
	// Step 1: Marshal resBody to JSON
	bodyJSON, err := json.Marshal(resBody)
	if err != nil {
		return err
	}

	// Step 2: Create an instance of the struct type that outputSchema represents
	schemaType := reflect.TypeOf(outputSchema)
	if schemaType.Kind() == reflect.Ptr {
		schemaType = schemaType.Elem()
	}
	schemaInstance := reflect.New(schemaType).Interface()

	// Unmarshal the JSON into the schema instance
	if err := json.Unmarshal(bodyJSON, schemaInstance); err != nil {
		return err
	}

	// Step 3: Validate the struct
	if err := validate.Struct(schemaInstance); err != nil {
		return err
	}

	return nil
}
