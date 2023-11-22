package main

import (
	bluerpc "github.com/S-IR/blueRPC/blueRPC"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Input struct {
	House string `query:"house" validate:"required"`
}

type Output struct {
	FieldOneOut   string `json:"fieldOneOut" validate:"required"`
	FieldTwoOut   string `json:"fieldTwoOut" validate:"required"`
	FieldThreeOut string `json:"fieldThreeOut" validate:"required"`
}

func main() {
	validate := validator.New(validator.WithRequiredStructEnabled())

	app := bluerpc.New(&bluerpc.Config{
		OutputPath:  "./some-file.ts",
		ValidatorFn: validate.Struct,
		FiberConfig: &fiber.Config{},
	})

	proc := bluerpc.NewQuery[Input, Output](app, func(ctx *fiber.Ctx, queryParams Input) (*bluerpc.Res[Output], error) {
		return &bluerpc.Res[Output]{
			Status: 200,
			Header: bluerpc.Header{},
			Body: Output{
				FieldOneOut:   "dwa",
				FieldTwoOut:   "dwadwa",
				FieldThreeOut: "dwadwadwa",
			},
		}, nil
	})

	users := app.Group("/users")
	proc.Attach(users, "/hello")

	app.Listen(":3000")

}

// func main() {
// 	app := bluerpc.New()

// 	type Output struct {
// 		Something string
// 	}

// 	helloGrp := app.Group("/hello")

// 	worldQuery := bluerpc.NewQuery[any, Output](app, func(ctx *fiber.Ctx, queryParams any) (*bluerpc.Res[Output], error) {
// 		return &bluerpc.Res[Output]{
// 			Status: 200,
// 			Body: Output{
// 				Something: "hello",
// 			},
// 		}, nil
// 	})
// 	worldQuery.Attach(helloGrp, "/world")

// 	app.Listen(":3000")
// }
