package bluerpc

import (
	"net/http"

	"github.com/go-playground/validator"
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

	method    Method
	handler   Handler
	validator *validator.Validate
}

type IDK interface {
	Use(args ...interface{})

	Query(Handler)
	Mutation(Handler)
}

func NewProcedure() *Procedure {
	return &Procedure{}
}
func (p *Procedure) Input(schema interface{}) *Procedure {
	p.inputSchema = schema

	if p.validator == nil {
		validator := validator.New()
		p.validator = validator
	}

	return p
}

func (p *Procedure) Output(schema interface{}) *Procedure {
	p.outputSchema = schema

	if p.validator == nil {
		validator := validator.New()
		p.validator = validator
	}

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
