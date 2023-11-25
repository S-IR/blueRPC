package bluerpc

import (
	"fmt"

	genTypescript "github.com/S-IR/blueRPC/blueRPC/genTS"
	"github.com/gofiber/fiber/v2"
)

func addQueryProcedure[T fiber.Router, queryParams any, input any, output any](handler T, basePath, slug string, proc *Procedure[queryParams, input, output]) {
	validatorFn := *proc.validatorFn
	FullHandler := func(c *fiber.Ctx) error {
		input, err := validateQuery(c, validatorFn, proc)
		if err != nil {
			return err
		}

		res, err := proc.queryHandler(c, input)

		if err != nil {
			return err
		}
		err = validateOutput(validatorFn, proc, res)

		if err != nil {
			return err
		}

		err = setHeaders(c, res.Header)
		if err != nil {
			return err
		}

		return c.JSON(res.Body)
	}

	if !proc.app.config.disableGenerateTS {

		params := *new(queryParams)

		output := *new(output)

		fullRoute := fmt.Sprintf("%s%s", basePath, slug)

		genTypescript.AddProcedureToTree(fullRoute, params, nil, output, genTypescript.Method(QUERY))

	}
	handler.Get(slug, FullHandler)

}

func addMutationProcedure[T fiber.Router, queryParams any, input any, output any](handler T, basePath, slug string, proc *Procedure[queryParams, input, output]) {

	validatorFn := *proc.validatorFn
	FullHandler := func(c *fiber.Ctx) error {

		queryParams, err := validateQuery(c, validatorFn, proc)
		if err != nil {
			return err
		}

		input, err := validateInput(c, validatorFn, proc)
		if err != nil {
			return err
		}

		res, err := proc.mutationHandler(c, queryParams, input)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		err = validateOutput(validatorFn, proc, res)
		if err != nil {
			return err
		}

		err = setHeaders(c, res.Header)
		if err != nil {
			return err
		}

		return c.JSON(res.Body)
	}
	if !proc.app.config.disableGenerateTS {
		params := *new(queryParams)

		input := *new(input)

		output := *new(output)

		fullRoute := fmt.Sprintf("%s%s", basePath, slug)
		genTypescript.AddProcedureToTree(fullRoute, params, input, output, genTypescript.Method(MUTATION))
	}
	handler.Post(slug, FullHandler)
}

func validateQuery[queryParams any, input any, output any](c *fiber.Ctx, validatorFn validatorFn, proc *Procedure[queryParams, input, output]) (queryParams, error) {
	queryParamInstance := new(queryParams)
	if proc.queryParamsSchema == nil || validatorFn == nil {
		return *queryParamInstance, nil
	}
	if err := c.QueryParser(queryParamInstance); err != nil {
		return *queryParamInstance, err
	}

	if err := validatorFn(queryParamInstance); err != nil {

		return *queryParamInstance, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}

	}

	return *queryParamInstance, nil

}
func validateInput[queryParams any, input any, output any](c *fiber.Ctx, validatorFn validatorFn, proc *Procedure[queryParams, input, output]) (input, error) {
	inputInstance := new(input)
	if proc.inputSchema == nil || validatorFn == nil {
		return *inputInstance, nil
	}

	if err := c.BodyParser(inputInstance); err != nil {
		fmt.Println("err here at bodyParser")
		return *inputInstance, err
	}
	// Validate the struct
	if err := validatorFn(inputInstance); err != nil {

		return *inputInstance, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}

	}
	return *inputInstance, nil
}
func validateOutput[queryParams any, input any, output any](validatorFn validatorFn, proc *Procedure[queryParams, input, output], res *Res[output]) error {
	if proc.outputSchema == nil || validatorFn == nil {
		return nil
	}

	if err := validatorFn(res.Body); err != nil {
		panic(err.Error())
	}
	return nil
}
func setHeaders(ctx *fiber.Ctx, header Header) error {
	if header.Authorization != "" {
		ctx.Set("Authorization", header.Authorization)
	}
	if header.CacheControl != "" {
		ctx.Set("Cache-Control", header.CacheControl)
	}
	if header.ContentEncoding != "" {
		ctx.Set("Content-Encoding", header.ContentEncoding)
	}
	if header.ContentType != "" {
		ctx.Set("Content-Type", header.ContentType)
	}
	if header.Expires != "" {
		ctx.Set("Expires", header.Expires)
	}
	for _, cookie := range header.Cookies {
		if cookie != nil {
			ctx.Cookie(cookie)
		}
	}
	return nil
}
