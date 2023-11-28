package bluerpc

import (
	"time"
)

// func TestTSPerformance(t *testing.T) {
// 	fmt.Printf(fiber.DefaultColors.Green + "TESTING PERFORMANCE: \n" + fiber.DefaultColors.Reset)

// 	validate := validator.New(validator.WithRequiredStructEnabled())
// 	app := New(&Config{
// 		OutputPath:  "./local-some-file.ts",
// 		ValidatorFn: validate.Struct,
// 		FiberConfig: &fiber.Config{},
// 	})

// 	query := NewQuery[test_query, test_output](app, func(ctx *fiber.Ctx, queryParams test_query) (*Res[test_output], error) {
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
// 	startCh := make(chan struct{})

// 	perfLoop := func(num int) time.Duration {
// 		currGroup := app.Group("/start")
// 		fmt.Printf("testing perf for %d", num)
// 		for i := 0; i < num; i++ {
// 			fmt.Printf("i %d \n", i)
// 			newGrp := currGroup.Group(fmt.Sprintf("depth%d", i))
// 			query.Attach(newGrp, "/query")
// 			mut.Attach(newGrp, "/mutation")

// 		}
// 		start := time.Now()

// 		var mut sync.Mutex

// 		go func() {
// 			err := app.Listen(":3000")
// 			if err != nil {
// 				fmt.Println("Error starting the server:", err)
// 				return
// 			}

// 			// Signal server readiness
// 			mut.Lock()
// 			defer mut.Unlock()
// 			close(startCh)
// 		}()

// 		// Wait for server readiness signal
// 		mut.Lock()
// 		defer mut.Unlock()
// 		<-startCh

// 		// Record elapsed time
// 		elapsed := time.Since(start)

// 		// Shut down the server
// 		app.Shutdown()

// 		return elapsed
// 	}

// 	avgTenTime := getAvg(func() time.Duration {
// 		// Replace the arguments with your actual arguments
// 		return perfLoop(10)
// 	})
// 	fmt.Printf(fiber.DefaultColors.Green+"AVERAGE TIME FOR GENERATING DEPTH OF 10: %s\n"+fiber.DefaultColors.Reset, avgTenTime)

// 	avgHundredTime := getAvg(func() time.Duration {
// 		return perfLoop(100)
// 	})
// 	fmt.Printf(fiber.DefaultColors.Green+"AVERAGE TIME FOR GENERATING DEPTH OF 100: %s \n, difference of by %s from 10\n"+fiber.DefaultColors.Reset, avgHundredTime, avgHundredTime-avgTenTime)

// 	avgHThousandTime := getAvg(func() time.Duration {
// 		return perfLoop(1000)
// 	})
// 	fmt.Printf(fiber.DefaultColors.Green+"AVERAGE TIME FOR GENERATING DEPTH OF 1000: %s \n, difference of by %s from 100\n"+fiber.DefaultColors.Reset, avgHThousandTime, avgHThousandTime-avgHundredTime)

// 	avgHTenThousandTime := getAvg(func() time.Duration {
// 		return perfLoop(10000)
// 	})
// 	fmt.Printf(fiber.DefaultColors.Green+"AVERAGE TIME FOR GENERATING DEPTH OF 10 000: %s \n, difference of by %s from 1000\n"+fiber.DefaultColors.Reset, avgHTenThousandTime, avgHTenThousandTime-avgHThousandTime)

// }

func getAvg(someFunc func() time.Duration) time.Duration {
	var total time.Duration

	// Run the function 100 times
	for i := 0; i < 100; i++ {
		// Measure the time taken by the function
		duration := someFunc()
		total += duration
	}

	// Calculate the average time
	avg := total / 100
	return avg

}
