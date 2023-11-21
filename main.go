package main

import (
	bluerpc "github.com/S-IR/blueRPC/blueRPC"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Input struct {
	FieldOne   string `json:"fieldOne" validate:"required"`
	FieldTwo   string `json:"fieldTwo" validate:"required"`
	FieldThree string `json:"fieldThree" validate:"required"`
}

type Output struct {
	FieldOneOut   string `json:"fieldOneOut" validate:"required"`
	FieldTwoOut   string `json:"fieldTwoOut" validate:"required"`
	FieldThreeOut string `json:"fieldThreeOut" validate:"required"`
}

func main() {
	validate := validator.New(validator.WithRequiredStructEnabled())

	app := bluerpc.New(&fiber.Config{}, &bluerpc.Config{
		OutputPath:  "./some-file.ts",
		ValidatorFn: validate.Struct,
	})

	proc := bluerpc.NewProcedure[Input, Output](app).Query(func(ctx *fiber.Ctx, input Input) (bluerpc.Res[Output], error) {
		return bluerpc.Res[Output]{
			Status: 200,
			Body: Output{
				FieldOneOut:   "dwa",
				FieldTwoOut:   "dwa",
				FieldThreeOut: "dwadwadwa",
			},
		}, nil
	})

	users := app.Group("/users")
	proc.Attach(users, "/hello")

	app.Listen(":3000")

}
