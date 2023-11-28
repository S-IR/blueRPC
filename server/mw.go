package bluerpc

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DefaultErrorMiddleware(c *fiber.Ctx) error {
	// Execute the next handler
	err := c.Next()
	fmt.Println("triggered error mw")
	fmt.Println("err", err)

	// Check if there was an error
	if err != nil {
		// This is a Fiber error type
		if e, ok := err.(*fiber.Error); ok {
			if e.Code >= 500 {
				return c.Status(e.Code).JSON(fiber.Map{
					"message": "An error has occurred. Please try again later",
				})
			}

			fmt.Println("reached here")

			return c.Status(e.Code).JSON(fiber.Map{
				"message": e.Message,
			})
		}

		// This handles any other type of error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// No error, continue with next middleware
	return nil
}
