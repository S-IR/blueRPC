package bluerpc

import "github.com/gofiber/fiber/v2"

type Res[T any] struct {
	Status int
	Header Header
	Body   interface{}
}

type Handler[input any, output any] func(ctx *fiber.Ctx, input input) (Res[output], error)
