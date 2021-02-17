package main

import (
	"net/http"
	"testing"
)

func TestGetRandomHttpRequests(t *testing.T) {
	setup()
	runtime := NewRuntimeInfo()
	runtime.HttpRequestMethodPercentage = map[string]int{"GET": 5, "PUT": 5}
	client := NewBenchmarkClient(runtime.TargetUrl, runtime.HttpMethods, runtime.HttpHeaders, runtime.HttpRequestBody, runtime.HttpRequestMethodPercentage)
	requests := client.GetRandomHttpRequests()

	if len(requests) != 10 {
		t.Fatal("Request is setting GET and PUT that request array is 10 length")
	}

	var putCnt, getCnt int
	for _, c := range client.GetRandomHttpRequests() {
		if c.Method == http.MethodGet {
			getCnt++
		}
		if c.Method == http.MethodPut {
			putCnt++
		}
		if c.Method != http.MethodGet && c.Method != http.MethodPut {
			t.Fatal("Method is always GET ot PUT")
		}
		if c.Header.Get("Authorization") != "Bearer xxxx" {
			t.Fatal("Header was setting Bearer xxxx")
		}
	}

	if getCnt != 5 {
		t.Fatal("GET request is 5 times")
	}
	if putCnt != 5 {
		t.Fatal("PUT request is 5 times")
	}
	clean()
}

func TestShuffleRequest(t *testing.T) {
	var requests []*http.Request
	for i := 0; i < 10; i++ {
		var request *http.Request
		if i < 5 {
			request, _ = http.NewRequest(http.MethodGet, "", nil)
		} else {
			request, _ = http.NewRequest(http.MethodPut, "", nil)
		}
		requests = append(requests, request)
	}

	// Check Request slice ordering when not shuffle
	for i, r := range requests {
		if i < 5 && r.Method != http.MethodGet {
			t.Fatal("GET when 5 or less")
		}

		if i > 5 && r.Method != http.MethodPut {
			t.Fatal("GET when 5 or more")
		}
	}

	shuffleRequest(requests)
}

func TestAttack(t *testing.T) {
	setup()
	durationSeconds = 3
	runtime := NewRuntimeInfo()
	runtime.HttpMethods = []string{http.MethodGet}
	runtime.HttpRequestMethodPercentage = map[string]int{"GET": 10}
	client := NewBenchmarkClient(runtime.TargetUrl, runtime.HttpMethods, runtime.HttpHeaders, runtime.HttpRequestBody, runtime.HttpRequestMethodPercentage)

	result := client.Attack(1)
	if len(result.Get) == 0 {
		t.Fatalf("Failed to Attack %d", result.Get)
	}
	clean()
}
