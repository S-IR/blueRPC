package main

import (
	bluerpc "github.com/S-IR/blueRPC/blueRPC"
	"github.com/gofiber/fiber/v2"
)

type Input struct {
	FieldOne   string `json:"fieldOne" validate:"required"`
	FieldTwo   string `json:"fieldTwo" validate:"required"`
	FieldThree string `json:"fieldThree" validate:"required"`
}

type Output struct {
	FieldOneOut      string `validate:"required"`
	FieldTwoOut      string `validate:"required"`
	FieldThreeStruct string `validate:"required"`
}

func main() {

	app := bluerpc.New(&fiber.Config{}, &bluerpc.Config{
		OutputPath: "./some-file.ts",
	})

	app.Group("/users", &bluerpc.ProcedureList{
		"hello": bluerpc.NewProcedure().Input(Input{}).Output(Output{}).Query(func(c *fiber.Ctx) (bluerpc.Res, error) {

			return bluerpc.Res{
				Status: 200,
				Body: map[string]string{
					"fieldOneOut":      "value1",
					"fieldTwoOutdd":    "value2",
					"fieldThreeStruct": "value3",
				},
			}, nil

		}),
	})
	app.Listen(":3000")

}
