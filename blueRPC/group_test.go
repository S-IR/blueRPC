package bluerpc

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func TestGroup(t *testing.T) {

	validate := validator.New(validator.WithRequiredStructEnabled())
	fmt.Println(fiber.DefaultColors.Green + "TESTING NESTED ROUTE" + fiber.DefaultColors.Reset)

	fmt.Println(fiber.DefaultColors.Green + "TESTING INVALID QUERY PARAMS")
	app := New(&Config{
		OutputPath:  "./some-file.ts",
		ValidatorFn: validate.Struct,
		FiberConfig: &fiber.Config{},
	})

	proc := NewQuery[any, test_output](app, func(ctx *fiber.Ctx, queryParams any) (*Res[test_output], error) {
		return &Res[test_output]{
			Status: 200,
			Header: Header{},
			Body: test_output{
				FieldOneOut:   "dwa",
				FieldTwoOut:   "dwadwa",
				FieldThreeOut: "dwadwadwa",
			},
		}, nil
	})
	depthOne := app.Group("/depth1")
	depthTwo := depthOne.Group("/depth2")

	proc.Attach(depthTwo, "/test")
	req, err := http.NewRequest("GET", "http://localhost:3000/bluerpc/depth1/depth2/test", nil)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not create a new request", err.Error())
	}
	res, err := app.Test(req)

	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not do the request", err.Error())
	}

	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not convert status code to number", err.Error())

	}
	if res.StatusCode > 300 {
		t.Fatalf(fiber.DefaultColors.Red+"Server did not respond with a 2xx status, actual status %d", res.StatusCode)

	}

	fmt.Println(fiber.DefaultColors.Green + "PASSED NESTED ROUTE TEST" + fiber.DefaultColors.Reset)

}
