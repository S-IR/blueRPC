package bluerpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type test_queryParams struct {
	Something string `query:"query" validate:"required"`
}

type test_input struct {
	House string `validate:"required"`
}

type test_output struct {
	FieldOneOut   string `json:"fieldOneOut" validate:"required"`
	FieldTwoOut   string `json:"fieldTwoOut" `
	FieldThreeOut string `json:"fieldThreeOut" validate:"required"`
}

func TestQuery(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	fmt.Println(fiber.DefaultColors.Green + "TESTING INVALID QUERY PARAMS")
	app := New(&Config{
		OutputPath:  "./some-file.ts",
		ValidatorFn: validate.Struct,
		FiberConfig: &fiber.Config{},
	})

	proc := NewQuery[test_queryParams, test_output](app, func(ctx *fiber.Ctx, queryParams test_queryParams) (*Res[test_output], error) {
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
	proc.Attach(app, "/test")

	// app.Listen(":3000")
	req, err := http.NewRequest("GET", "http://localhost:3000/test", nil)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not create a new request", err.Error())
	}
	res, err := app.Test(req)

	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not do the request", err.Error())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not read the body", err.Error())
	}
	type DefaultResError struct {
		Message string `json:"message"`
	}
	var resError DefaultResError
	if err := json.Unmarshal(body, &resError); err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Failed to unmarshal response: %v", err)
	}
	if resError.Message == "" {
		t.Fatalf(fiber.DefaultColors.Red + "The server responded without an error")
	}

	fmt.Println(fiber.DefaultColors.Green + "PASSED INVALID QUERY")

	// TESTING VALID QUERY PARAMS
	fmt.Println(fiber.DefaultColors.Green + "TESTING VALID QUERY PARAMS")
	req, err = http.NewRequest("GET", "http://localhost:3000/test?query=dwa", nil)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not create a new request", err.Error())
	}
	res, err = app.Test(req)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not do the request", err.Error())
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not read the body", err.Error())
	}

	var output test_output
	if err := json.Unmarshal(body, &output); err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Failed to unmarshal response: %v", err)
	}

	if output.FieldOneOut == "" || output.FieldTwoOut == "" || output.FieldThreeOut == "" {
		t.Fatalf(fiber.DefaultColors.Red+"The server responded with an invalid output response", string(body))
	}

	fmt.Println(fiber.DefaultColors.Green + "PASSED VALID QUERY")
}

func TestMutation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	fmt.Println(fiber.DefaultColors.Green + "TESTING INVALID QUERY PARAMS")
	app := New(&Config{
		OutputPath:  "./some-file.ts",
		ValidatorFn: validate.Struct,
		FiberConfig: &fiber.Config{},
	})

	proc := NewMutation[test_queryParams, test_input, test_output](app, func(ctx *fiber.Ctx, queryParams test_queryParams, input test_input) (*Res[test_output], error) {

		return &Res[test_output]{
			Status: 200,
			Body: test_output{
				FieldOneOut:   "dwaawdwa",
				FieldTwoOut:   "dwa",
				FieldThreeOut: "dawdwadwadwa",
			},
		}, nil

	})
	proc.Attach(app, "/test")
	// app.Listen(":3000")

	inputData := test_input{
		House: "hello world",
	}

	jsonData, err := json.Marshal(inputData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Log the JSON payload before sending the request

	req, err := http.NewRequest("POST", "http://localhost:3000/test", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	res, err := app.Test(req)

	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not do the request", err.Error())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not read the body", err.Error())
	}
	type DefaultResError struct {
		Message string
	}
	var resError DefaultResError
	if err := json.Unmarshal(body, &resError); err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Failed to unmarshal response: %v", err)
	}
	if resError.Message == "" {
		t.Fatalf(fiber.DefaultColors.Red+"The server responded without an error, %s", string(body))
	}

	fmt.Println(fiber.DefaultColors.Green + "PASSED INVALID MUTATION")

	// TESTING VALID QUERY PARAMS

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(fiber.DefaultColors.Green + "TESTING VALID MUTATION PARAMS")
	req, err = http.NewRequest("POST", "http://localhost:3000/test?query=dwa", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not create a new request", err.Error())
	}
	res, err = app.Test(req)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not do the request", err.Error())
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not read the body", err.Error())
	}

	var output test_output
	if err := json.Unmarshal(body, &output); err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Failed to unmarshal response: %v", err)
	}

	if output.FieldOneOut == "" || output.FieldTwoOut == "" || output.FieldThreeOut == "" {
		t.Fatalf(fiber.DefaultColors.Red+"The server responded with an invalid output response, %s", string(body))
	}

	fmt.Println(fiber.DefaultColors.Green + "PASSED VALID MUTATION")
	fmt.Println(fiber.DefaultColors.Green + "TESTING INVALID OUTPUT")

	fakeProc := NewMutation[test_queryParams, test_input, test_output](app, func(ctx *fiber.Ctx, queryParams test_queryParams, input test_input) (*Res[test_output], error) {

		return &Res[test_output]{
			Status: 200,
			Body: test_output{
				FieldOneOut:   "",
				FieldTwoOut:   "dwa",
				FieldThreeOut: queryParams.Something,
			},
		}, nil

	})
	fakeProc.Attach(app, "/error")

	req, err = http.NewRequest("POST", "http://localhost:3000/error?query=dwa", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not create a new request", err.Error())
	}
	res, err = app.Test(req)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not do the request", err.Error())
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Could not read the body", err.Error())
	}

	var err_output DefaultResError
	if err := json.Unmarshal(body, &err_output); err != nil {
		t.Fatalf(fiber.DefaultColors.Red+"Failed to unmarshal response: %v", err)
	}

	if err_output.Message == "" {
		t.Fatalf(fiber.DefaultColors.Red+"The body output error response is not proper : %v", err)

	}

	fmt.Println(fiber.DefaultColors.Green + "PASSED INVALID OUTPUT")

}
