package bluerpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

type test_query struct {
	Something string `query:"query" validate:"required"`
}

type test_input struct {
	House string `json:"house" validate:"required"`
}

type test_output struct {
	FieldOneOut   string `json:"fieldOneOut" validate:"required"`
	FieldTwoOut   string `json:"fieldTwoOut" `
	FieldThreeOut string `json:"fieldThreeOut" validate:"required"`
}

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
	resp, err := http.Get("http://localhost:3000/bluerpc/hello/world")
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
