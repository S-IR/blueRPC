package bluerpc

import (
	"reflect"
)

type Method string

var (
	QUERY    Method = "query"
	MUTATION Method = "mutation"
)

type Procedure[queryParams any, input any, output any] struct {
	queryParamsSchema *queryParams
	inputSchema       *input
	outputSchema      *output

	method      Method
	app         *App
	validatorFn *validatorFn

	queryHandler    Query[queryParams, output]
	mutationHandler Mutation[queryParams, input, output]
}

// Creates a new query procedure that can be attached to groups / app root.
// The generic arguments specify the structure for validating query parameters (the query Params, the body of the request and the resulting handler output).
// Use any to avoid validation
func NewMutation[queryParams any, input any, output any](app *App, mutation Mutation[queryParams, input, output]) *Procedure[queryParams, input, output] {

	var inputInstance *input
	var outputInstance *output
	if reflect.TypeOf(new(input)).Elem().Kind() == reflect.Struct {
		temp := new(input)
		inputInstance = temp
	}

	// Check if output is a struct
	if reflect.TypeOf(new(output)).Elem().Kind() == reflect.Struct {
		temp := new(output)
		outputInstance = temp
	}

	return &Procedure[queryParams, input, output]{
		app:             app,
		validatorFn:     &app.config.ValidatorFn,
		inputSchema:     inputInstance,
		outputSchema:    outputInstance,
		method:          MUTATION,
		mutationHandler: mutation,
	}
}

// Creates a new query procedure that can be attached to groups / app root.
// The generic arguments specify the structure for validating query parameters (the query Params and the resulting handler output).
// Use any to avoid validation
func NewQuery[queryParams any, output any](app *App, query Query[queryParams, output]) *Procedure[queryParams, any, output] {

	var queryParamsInstance *queryParams
	var outputInstance *output
	if reflect.TypeOf(new(queryParams)).Elem().Kind() == reflect.Struct {
		temp := new(queryParams)
		queryParamsInstance = temp
	}

	// Check if output is a struct
	if reflect.TypeOf(new(output)).Elem().Kind() == reflect.Struct {
		temp := new(output)
		outputInstance = temp
	}

	return &Procedure[queryParams, any, output]{
		app:               app,
		validatorFn:       &app.config.ValidatorFn,
		outputSchema:      outputInstance,
		queryParamsSchema: queryParamsInstance,
		method:            QUERY,
		queryHandler:      query,
	}
}

// Attaches the given Procedure to a group
func (p *Procedure[queryParams, input, output]) Attach(grp ValidRouter, path string) {
	router := grp.getFiberRouter()

	switch p.method {
	case QUERY:
		addQueryProcedure(router, path, p)
	case MUTATION:
		addMutationProcedure(router, path, p)
	}
}
