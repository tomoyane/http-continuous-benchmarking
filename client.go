package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"math/rand"
	"net/http"
)

const (
	durationSeconds = 15
	requestDuration = durationSeconds * time.Second
)

type BenchmarkClient interface {
	Attack() Result
}

type HttpClient struct {
	HttpClient         *http.Client
	RandomHttpRequests []*http.Request
}

type Result struct {
	Get    []int
	Post   []int
	Put    []int
	Patch  []int
	Delete []int
}

// New BenchmarkClient
func NewBenchmarkClient(url string, methods []string, headers map[string]string, body io.Reader, percentages map[string]int) BenchmarkClient {
	var requests []*http.Request
	for _, method := range methods {
		var request *http.Request
		for targetMethod, percentage := range percentages {
			if strings.EqualFold(method, targetMethod) {
				// Generate request per percentage method
				for i := 0; i < percentage; i++ {
					if !strings.EqualFold(method, http.MethodGet) {
						request, _ = http.NewRequest(method, url, body)
					} else {
						request, _ = http.NewRequest(method, url, nil)
					}
					// Set headers
					for headerKey, headerValue := range headers {
						request.Header.Set(headerKey, headerValue)
					}
					requests = append(requests, request)
				}
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(requests), func(i, j int) { requests[i], requests[j] = requests[j], requests[i] })
	fmt.Printf("Request info %v\n", requests)

	client := new(http.Client)
	return HttpClient{
		HttpClient:         client,
		RandomHttpRequests: requests,
	}
}

func (h HttpClient) Attack() Result {
	var getLatency, postLatency, putLatency, patchLatency, deleteLatency []int
	fmt.Printf("#### Start benchmark duration %d\n", requestDuration)
	for begin := time.Now(); time.Since(begin) < requestDuration; {
		// Random Http Method request
		for _, request := range h.RandomHttpRequests {
			start := time.Second / time.Millisecond
			res, err := h.HttpClient.Do(request)

			if err == nil && (res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated) {
				end := time.Second / time.Millisecond
				latency := end - start
				switch request.Method {
				case http.MethodGet:
					getLatency = append(getLatency, int(latency))
				case http.MethodPost:
					postLatency = append(postLatency, int(latency))
				case http.MethodPut:
					putLatency = append(putLatency, int(latency))
				case http.MethodPatch:
					patchLatency = append(patchLatency, int(latency))
				case http.MethodDelete:
					deleteLatency = append(deleteLatency, int(latency))
				}
			}
		}
	}
	fmt.Println("#### End benchmark")
	return Result{
		Get:    getLatency,
		Post:   postLatency,
		Put:    putLatency,
		Patch:  patchLatency,
		Delete: deleteLatency,
	}
}
