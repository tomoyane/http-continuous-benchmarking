package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"math/rand"
	"net/http"
)

var durationSeconds = 5
var warmupSeconds = 5

type BenchmarkClient interface {
	Attack(attackNum int) Result

	Warmup()

	GetRandomHttpRequests() []*http.Request
}

type HttpClient struct {
	HttpClient         *http.Client
	RandomHttpRequests []*http.Request
	RequestDuration    time.Duration
}

type Result struct {
	Get     []int
	Post    []int
	Put     []int
	Patch   []int
	Delete  []int
	ErrData map[string]map[int]int
}

// NewBenchmarkClient New BenchmarkClient
func NewBenchmarkClient(url string, methods []string, headers map[string]string, body io.Reader, percentages map[string]int) BenchmarkClient {
	var requests []*http.Request
	for _, method := range methods {
		var request *http.Request
		for targetMethod, percentage := range percentages {
			if strings.EqualFold(method, targetMethod) {
				// GenerateCharts request per percentage method
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

	shuffleRequest(requests)
	fmt.Print("HTTP request pattern according to the ratio = ")
	for _, r := range requests {
		fmt.Printf("%s ", r.Method)
	}
	fmt.Println()
	client := new(http.Client)
	return HttpClient{
		HttpClient:         client,
		RandomHttpRequests: requests,
		RequestDuration:    time.Duration(durationSeconds) * time.Second,
	}
}

func (h HttpClient) Attack(attackNum int) Result {
	var getLatency, postLatency, putLatency, patchLatency, deleteLatency []int
	errData := make(map[string]map[int]int)
	fmt.Printf("(Thread-%d): Start attack for duration %d seconds\n", attackNum, durationSeconds)
	for begin := time.Now(); time.Since(begin) < h.RequestDuration; {
		// Random Http Method request
		for _, request := range h.RandomHttpRequests {
			start := makeTimestamp()
			res, err := h.HttpClient.Do(request)
			if err != nil {
			} else if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
				end := makeTimestamp()
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
			} else {
				if _, ok := errData[request.Method]; ok {
					errData[request.Method][res.StatusCode] = errData[request.Method][res.StatusCode] + 1
				} else {
					errData[request.Method] = map[int]int{res.StatusCode: 1}
				}
			}
			defer func() {
				if res != nil {
					io.Copy(ioutil.Discard, res.Body)
					res.Body.Close()
				}
			}()
		}
	}
	fmt.Printf("(Thread-%d): End attack \n", attackNum)
	return Result{
		Get:     getLatency,
		Post:    postLatency,
		Put:     putLatency,
		Patch:   patchLatency,
		Delete:  deleteLatency,
		ErrData: errData,
	}
}

func (h HttpClient) Warmup() {
	fmt.Printf("Start warnmup for duration %d seconds\n", warmupSeconds)
	for begin := time.Now(); time.Since(begin) < h.RequestDuration; {
		// Random Http Method request
		for _, request := range h.RandomHttpRequests {
			h.HttpClient.Do(request)
		}
	}
	fmt.Println("End warmup")
	fmt.Println()
}

func (h HttpClient) GetRandomHttpRequests() []*http.Request {
	return h.RandomHttpRequests
}

func shuffleRequest(requests []*http.Request) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(requests), func(i, j int) { requests[i], requests[j] = requests[j], requests[i] })
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
