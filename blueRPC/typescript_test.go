package bluerpc

//This is a bad test. It needs to be modified

// const outputPath = "./test-file.ts"
// const output = "type bluerpc ={depth1:{depth2:{test:{query:(queryParams:{ Something: string,})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});mutation:(input : {queryParams:{ Something: string,},input:{ house: string,}})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});}},zap:{mutation:(input : {queryParams:{ Something: string,},input:{ house: string,}})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});}},helloWorld:{query:(queryParams:{ Something: string,})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});}}"

// func TestTypescriptGen(t *testing.T) {
// 	fmt.Println(fiber.DefaultColors.Green + "TESTING TYPESCRIPT GENERATION" + fiber.DefaultColors.Reset)

// 	validate := validator.New(validator.WithRequiredStructEnabled())
// 	app := New(&Config{
// 		OutputPath:  "./some-file.ts",
// 		ValidatorFn: validate.Struct,
// 		FiberConfig: &fiber.Config{},
// 	})

// 	proc := NewQuery[test_query, test_output](app, func(ctx *fiber.Ctx, queryParams test_query) (*Res[test_output], error) {
// 		return &Res[test_output]{
// 			Status: 200,
// 			Body: test_output{
// 				FieldOneOut:   "dwa",
// 				FieldTwoOut:   "dwadwa",
// 				FieldThreeOut: "dwadwadwa",
// 			},
// 		}, nil
// 	})
// 	mut := NewMutation[test_query, test_input, test_output](app, func(ctx *fiber.Ctx, queryParams test_query, input test_input) (*Res[test_output], error) {
// 		return &Res[test_output]{
// 			Status: 200,
// 			Body: test_output{
// 				FieldOneOut:   "dwadwa",
// 				FieldTwoOut:   "dwadwadwa",
// 				FieldThreeOut: "dwadwadwad",
// 			},
// 		}, nil
// 	})

// 	depthOne := app.Group("/depth1")
// 	depthTwo := depthOne.Group("/depth2")

// 	proc.Attach(depthTwo, "/test")
// 	mut.Attach(depthTwo, "/test")

// 	mut.Attach(depthOne, "/zap")

// 	proc.Attach(app, "/helloWorld")

// 	app.Listen(":3000")
// 	file, err := os.ReadFile(outputPath)

// 	if err != nil {
// 		t.Fatalf(fiber.DefaultColors.Red + "Could not read the output file" + err.Error())
// 	}
// 	if strings.TrimSpace(string(file)) != strings.TrimSpace(output) {
// 		t.Fatalf(fiber.DefaultColors.Red+"output of typescript is not proper,\n % s \n VS \n %s", string(file), output)
// 	}

// 	fmt.Println(fiber.DefaultColors.Green + "PASSED TYPESCRIPT GENERATION" + fiber.DefaultColors.Reset)
// 	err = os.Remove(output)
// 	if err != nil {
// 		panic(err)
// 	}

// }
