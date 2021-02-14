package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	// Required
	// ex: https://example.com
	EnvTargetUrl = "TARGET_URL"

	// Required
	// Comma separated
	// ex: PUT,GET
	EnvHttpMethods = "HTTP_METHODS"

	// Required
	// HashMap data structure
	// ex: {"Authorization": "Bearer ", "Content-Type": "application/json"}
	EnvHttpHeaders = "HTTP_HEADERS"

	// Required
	// Maximum is 3.
	// ex: 2
	EnvThreadNum = "THREAD_NUM"

	// Required
	// Maximum is 20. Takes up to 5 minutes
	// ex: 20
	EnvTrialNum = "TRIAL_NUM"

	// Optional
	// HashMap data structure
	// If only one http method, always 100 percent set method
	// ex: {"POST": 4, "GET": 6}
	EnvReqHttpMethodPercentages = "REQ_HTTP_METHOD_PERCENTAGES"

	// Optional
	// Using GitHub pages
	// ex: true || false
	EnvPermanent = "PERMANENT"

	// Optional
	// If not empty, always use body when not GET method
	// ex: {"email": "test@gmail.com", "password": "A_test12345-"}
	EnvHttpRequestBody = "HTTP_REQ_BODY"

	// Optional
	// ex: https://slack.com
	EnvSlackWebHookUrl = "SLACK_WEB_HOOK_URL"

	// Optional
	// If set this one, notify slack when do not achieve
	// ex: 200
	EnvSlackNotifyThreshHoldLatencyMillis = "SLACK_NOTIFY_THRESHOLD_LATENCY_MILLIS"
)

func main() {
	errMsg := ValidateEnv()
	if errMsg != nil {
		for _, v := range errMsg {
			fmt.Println(v)
		}
		return
	}

	runtime := NewRuntimeInfo()
	client := NewBenchmarkClient(
		runtime.TargetUrl,
		runtime.HttpMethods,
		runtime.HttpHeaders,
		runtime.HttpRequestBody,
		runtime.HttpRequestMethodPercentage,
	)
	calculator := NewCalculator(runtime.TrialNum)

	startTime := time.Now().UTC()
	for i := 1; i <= runtime.TrialNum; i++ {
		var wg sync.WaitGroup
		var result Result
		for i := 0; i < runtime.ThreadNum; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				data := client.Attack()
				result.Get = append(result.Get, data.Get...)
				result.Post = append(result.Post, data.Post...)
				result.Put = append(result.Put, data.Put...)
				result.Patch = append(result.Patch, data.Patch...)
				result.Delete = append(result.Delete, data.Delete...)
			}(i)
		}
		wg.Wait()

		calculator.CalculatePerTrial(result.Get, http.MethodGet, i)
		calculator.CalculatePerTrial(result.Post, http.MethodPost, i)
		calculator.CalculatePerTrial(result.Put, http.MethodPut, i)
		calculator.CalculatePerTrial(result.Patch, http.MethodPatch, i)
		calculator.CalculatePerTrial(result.Delete, http.MethodDelete, i)

		time.Sleep(1 * time.Second)
	}
	endTime := time.Now().UTC()

	metrics := calculator.GetMetricsResult()
	graph := NewGraph(metrics)
	charts := graph.Generate(metrics.TimeRange)
	graph.Output(charts, startTime, endTime)
}
