package main

import (
	bluerpc "github.com/S-IR/blueRPC/blueRPC"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Query struct {
	Something string `json:"query" validate:"required"`
}

type Input struct {
	House string `input:"house" validate:"required"`
}

type Output struct {
	FieldOneOut   string `json:"fieldOneOut" validate:"required"`
	FieldTwoOut   string `json:"fieldTwoOut" `
	FieldThreeOut string `json:"fieldThreeOut" validate:"required"`
}

func main() {
	validate := validator.New(validator.WithRequiredStructEnabled())

	app := bluerpc.New(&bluerpc.Config{
		OutputPath:  "./some-file.ts",
		ValidatorFn: validate.Struct,
		FiberConfig: &fiber.Config{},
	})

	proc := bluerpc.NewQuery[Query, Output](app, func(ctx *fiber.Ctx, queryParams Query) (*bluerpc.Res[Output], error) {
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

	mut := bluerpc.NewMutation[Query, Input, Output](app, func(ctx *fiber.Ctx, queryParams Query, input Input) (*bluerpc.Res[Output], error) {
		return &bluerpc.Res[Output]{
			Status: 200,
			Body: Output{
				FieldOneOut:   "dwa",
				FieldTwoOut:   "dwadwa",
				FieldThreeOut: "dwadwa",
			},
		}, nil
	})

	users := app.Group("/users")

	// currUser := users
	// for i := 0; i < 1000; i++ {
	// 	deeperUser := currUser.Group(fmt.Sprintf("/depth%d", i))

	// 	proc.Attach(deeperUser, "/hello")
	// 	mut.Attach(deeperUser, "/hello")
	// 	proc.Attach(deeperUser, "/bye")
	// 	mut.Attach(deeperUser, "/")

	// 	currUser = deeperUser

	// }

	deeperUser := users.Group("/depth1")
	deeperUserTwo := deeperUser.Group("/depth2")
	deeperUserThree := deeperUserTwo.Group("/depth3")

	proc.Attach(users, "/hello")
	proc.Attach(deeperUser, "/hello")
	proc.Attach(deeperUserTwo, "/hello")
	proc.Attach(deeperUserThree, "/hello")

	proc.Attach(users, "/bye")
	proc.Attach(deeperUser, "/bye")
	proc.Attach(deeperUserTwo, "/bye")
	proc.Attach(deeperUserThree, "/bye")

	mut.Attach(users, "/hello")
	mut.Attach(deeperUser, "/hello")
	mut.Attach(deeperUserTwo, "/hello")
	mut.Attach(deeperUserThree, "/hello")

	mut.Attach(users, "/bye")
	mut.Attach(deeperUser, "/bye")
	mut.Attach(deeperUserTwo, "/bye")
	mut.Attach(deeperUserThree, "/bye")

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
