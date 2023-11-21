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

	app.Group("/users", &bluerpc.ProcedureList{
		"hello": app.NewProcedure().Input(&Input{}).Query(func(c *fiber.Ctx) (bluerpc.Res[Output], error) {

			return bluerpc.Res[Output]{
				Status: 200,
				Body: Output{
					FieldTwoOut:   "value2",
					FieldThreeOut: "value3",
				},
			}, nil

		}),
	})
	app.Listen(":3000")

}
