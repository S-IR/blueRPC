package bluerpc

import (
	"errors"
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

type Procedure struct {
	inputSchema  interface{}
	outputSchema interface{}

	method      Method
	handler     Handler
	app         *App
	validatorFn *validatorFn
}

type IDK interface {
	Use(args ...interface{})

	Query(Handler)
	Mutation(Handler)
}

func (a *App) NewProcedure() *Procedure {
	return &Procedure{
		app:         a,
		validatorFn: &a.config.ValidatorFn,
	}
}

func (p *Procedure) Input(schema interface{}) *Procedure {
	if reflect.ValueOf(schema).Kind() != reflect.Ptr {
		panic(errors.New("Output schema must be a pointer to a struct"))
	}
	// Further, check if it's a pointer to a struct
	if reflect.Indirect(reflect.ValueOf(schema)).Kind() != reflect.Struct {
		panic(errors.New("Output schema must be a pointer to a struct"))
	}

	p.inputSchema = schema

	return p
}

func (p *Procedure) Output(schema interface{}) *Procedure {
	if reflect.ValueOf(schema).Kind() != reflect.Ptr {
		panic(errors.New("Output schema must be a pointer to a struct"))
	}
	// Further, check if it's a pointer to a struct
	if reflect.Indirect(reflect.ValueOf(schema)).Kind() != reflect.Struct {
		panic(errors.New("Output schema must be a pointer to a struct"))
	}

	p.outputSchema = schema

	return p
}

func (p *Procedure) Query(h Handler) *Procedure {
	p.method = QUERY
	p.handler = h
	return p
}

// func (p *Procedure) Attach(g *Group, endpoint string) error {

// 	if len(endpoint) == 0 {
// 		return fmt.Errorf("your endpoint must be at least 1 character long. Current value :%s", endpoint)
// 	}
// 	if endpoint[0] != '/' {
// 		return fmt.Errorf("your endpoint must start with a forward slash ( / ). Current value : %s", endpoint)
// 	}
// 	addProcedure(endpoint, p)
// 	return nil
// }
