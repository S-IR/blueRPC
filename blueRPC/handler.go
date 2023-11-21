package bluerpc

import "github.com/gofiber/fiber/v2"

type Res[T any] struct {
	Status int
	Header Header
	Body   interface{}
}
type Handler[T any] func(*fiber.Ctx) (Res[T], error)
