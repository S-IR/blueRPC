package bluerpc

import (
	"net/http"
	"reflect"
)

type Method string

var (
	QUERY    Method = "query"
	MUTATION Method = "mutation"
)

type Header struct {
	Authorization   string         // Credentials for authenticating the client to the server
	CacheControl    string         // Directives for caching mechanisms in both requests and responses
	ContentEncoding string         // The encoding of the body
	ContentType     string         // The MIME type of the body of the request (used with POST and PUT requests)
	Expires         string         // Gives the date/time after which the response is considered stale
	Cookies         []*http.Cookie //Cookies
}

type Procedure[input any, output any] struct {
	inputSchema  *input
	outputSchema *output

	method      Method
	handler     Handler[input, output]
	app         *App
	validatorFn *validatorFn
}

// Creates a new procedure to use around the app
// First generic argument is Input, Second is Output
func NewProcedure[input any, output any](app *App) *Procedure[input, output] {

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

	return &Procedure[input, output]{
		app:          app,
		validatorFn:  &app.config.ValidatorFn,
		inputSchema:  inputInstance,
		outputSchema: outputInstance,
	}
}

func (p *Procedure[input, output]) Query(h Handler[input, output]) *Procedure[input, output] {
	p.method = QUERY
	p.handler = h
	return p
}

func (p *Procedure[input, output]) Mutate(h Handler[input, output]) *Procedure[input, output] {
	p.method = MUTATION
	p.handler = h
	return p
}
func (p *Procedure[input, output]) Attach(grp *Group, path string) {
	addProcedure(grp.fiberRouter, path, p)

}
