package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	testTargetUrl       = "https://example.com"
	testHttpHeaders     = `{"Authorization": "Bearer xxxx"}`
	testThreadNum       = "10"
	testLoadTimeSeconds = "100"
	testPercentage      = `{"GET":5, "PUT":5}`
	testPermanent       = "false"
	testBody            = `{"test_key": "test_value"}`
	testSlackUrl        = "https://slack.com"
	testSlackThreshold  = "200"
)

func setup() {
	os.Setenv(EnvTargetUrl, testTargetUrl)
	os.Setenv(EnvHttpHeaders, testHttpHeaders)
	os.Setenv(EnvThreadNum, testThreadNum)
	os.Setenv(EnvTrialNum, testLoadTimeSeconds)
	os.Setenv(EnvReqHttpMethodRatio, testPercentage)
	os.Setenv(EnvPermanent, testPermanent)
	os.Setenv(EnvHttpRequestBody, testBody)
	os.Setenv(EnvSlackWebHookUrl, testSlackUrl)
	os.Setenv(EnvSlackNotifyThreshHoldLatencyMillis, testSlackThreshold)
}

func clean() {
	os.Setenv(EnvTargetUrl, "")
	os.Setenv(EnvHttpHeaders, "")
	os.Setenv(EnvThreadNum, "")
	os.Setenv(EnvTrialNum, "")
	os.Setenv(EnvReqHttpMethodRatio, "")
	os.Setenv(EnvPermanent, "")
	os.Setenv(EnvHttpRequestBody, "")
	os.Setenv(EnvSlackWebHookUrl, "")
	os.Setenv(EnvSlackNotifyThreshHoldLatencyMillis, "")
}

func TestNewRuntimeInfo(t *testing.T) {
	setup()
	runtime := NewRuntimeInfo()
	if runtime.TargetUrl != testTargetUrl {
		t.Fatalf("%s is not matched", EnvTargetUrl)
	}

	for _, v := range runtime.HttpHeaders {
		if !strings.Contains(testHttpHeaders, v) {
			t.Fatalf("%s is not matched", EnvHttpHeaders)
		}
	}

	expectThreadNum, _ := strconv.Atoi(testThreadNum)
	if runtime.ThreadNum != expectThreadNum {
		t.Fatalf("%s is not matched", EnvThreadNum)
	}

	expectLoadTime, _ := strconv.Atoi(testLoadTimeSeconds)
	if runtime.TrialNum != expectLoadTime {
		t.Fatalf("%s is not matched", EnvTrialNum)
	}

	for _, v := range runtime.HttpRequestMethodRatio {
		percent := strconv.Itoa(v)
		if !strings.Contains(testPercentage, percent) {
			t.Fatalf("%s is not matched", EnvReqHttpMethodRatio)
		}
	}

	for _, v := range runtime.HttpMethods {
		if !strings.Contains("GET,PUT", v) {
			t.Fatalf("Http Method is not matched")
		}
	}

	if strconv.FormatBool(runtime.Permanent) != testPermanent {
		t.Fatalf("%s is not matched", EnvPermanent)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(runtime.HttpRequestBody)
	if buf.String() != testBody {
		t.Fatalf("%s is not matched", EnvHttpRequestBody)
	}

	if runtime.SlackWebHookUrl != testSlackUrl {
		t.Fatalf("%s is not matched", EnvSlackWebHookUrl)
	}

	expectSlackThreadNum, _ := strconv.Atoi(testSlackThreshold)
	if runtime.SlackNotifyThreshHoldLatencyMillis != expectSlackThreadNum {
		t.Fatalf("%s is not matched", EnvSlackNotifyThreshHoldLatencyMillis)
	}

	clean()
}
