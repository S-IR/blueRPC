package bluerpc

import "github.com/gofiber/fiber/v2"

type Res struct {
	Status int
	Header Header
	Body   interface{}
}
type Handler func(*fiber.Ctx) (Res, error)
