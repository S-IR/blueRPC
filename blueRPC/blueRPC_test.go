package bluerpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func TestStart(t *testing.T) {
	app := New()

	type Output struct {
		Something string
	}

	helloGrp := app.Group("/hello")

	worldQuery := NewQuery[any, Output](app, func(ctx *fiber.Ctx, queryParams any) (*Res[Output], error) {
		return &Res[Output]{
			Body: Output{
				Something: "world",
			},
		}, nil
	})
	worldQuery.Attach(helloGrp, "/world")
	go func() {
		if err := app.Listen(":3000"); err != nil {
			t.Logf("Server failed to start: %v", err)
		}
	}()

	// Wait a bit for the server to start
	time.Sleep(time.Second * 1)

	// Make the request to the server
	resp, err := http.Get("http://localhost:3000/hello/world")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal and assert the response
	var output Output
	if err := json.Unmarshal(body, &output); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if output.Something != "world" {
		t.Errorf("Expected 'world', got '%s'", output.Something)
	}
	fmt.Println("PASSED TESTING START")

}

func TestInvalidInput(t *testing.T) {

	validator := validator.New(validator.WithRequiredStructEnabled())

	app := New(&Config{
		ValidatorFn: validator.Struct,
	})
	grp := app.Group("/hey")
	type Input struct {
		Something string `validate:"required"`
	}

	worldQuery := NewQuery[any, any](app, func(ctx *fiber.Ctx, queryParams any) (*Res[any], error) {
		return &Res[any]{
			Status: 200,
			Body: map[string]string{
				"hello": "world",
			},
		}, nil
	})
	worldQuery.Attach(grp, "/hey")
	go func() {
		if err := app.Listen(":3000"); err != nil {
			t.Logf("Server failed to start: %v", err)
		}
	}()

	// Wait a bit for the server to start
	time.Sleep(time.Second * 1)

	// Make the request to the server
	resp, err := http.Get("http://localhost:3000/hello/world")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	type DefaultResError struct {
		Message string
	}
	// Unmarshal and assert the response
	var resError DefaultResError
	if err := json.Unmarshal(body, &resError); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	fmt.Println("PASSED TESTING INVALID INPUT")

}
