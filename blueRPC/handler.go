package bluerpc

import (
	"github.com/gofiber/fiber/v2"
)

type Res[T any] struct {
	Status int
	Header Header
	Body   T
}
type Header struct {
	Authorization   string          // Credentials for authenticating the client to the server
	CacheControl    string          // Directives for caching mechanisms in both requests and responses
	ContentEncoding string          // The encoding of the body
	ContentType     string          // The MIME type of the body of the request (used with POST and PUT requests)
	Expires         string          // Gives the date/time after which the response is considered stale
	Cookies         []*fiber.Cookie //Cookies
}

// First Generic argument is QUERY PARAMS.
// Second is OUTPUT
type Query[queryParams any, output any] func(ctx *fiber.Ctx, queryParams queryParams) (*Res[output], error)

// First Generic argument is QUERY PARAMETERS.
// Second is INPUT.
// Third is OUTPUT.
type Mutation[queryParams any, input any, output any] func(ctx *fiber.Ctx, queryParams queryParams, input input) (*Res[output], error)
