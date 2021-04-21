package main

import (
	"io"
	"os"
	"strconv"
	"strings"

	"encoding/json"
)

type RuntimeInfo struct {
	TargetUrl                          string
	HttpMethods                        []string
	HttpHeaders                        map[string]string
	ThreadNum                          int
	TrialNum                           int
	HttpRequestMethodRatio             map[string]int
	Permanent                          bool
	HttpRequestBody                    io.Reader
	SlackWebHookUrl                    string
	SlackNotifyThreshHoldLatencyMillis int
}

// NewRuntimeInfo New RuntimeInfo from environment variable
func NewRuntimeInfo() RuntimeInfo {
	targetUrl := os.Getenv(EnvTargetUrl)
	headers := make(map[string]string)
	json.Unmarshal([]byte(os.Getenv(EnvHttpHeaders)), &headers)
	threadNum, _ := strconv.Atoi(os.Getenv(EnvThreadNum))
	trialNum, _ := strconv.Atoi(os.Getenv(EnvTrialNum))
	requestMethodRatio := make(map[string]int)
	json.Unmarshal([]byte(os.Getenv(EnvReqHttpMethodRatio)), &requestMethodRatio)
	var methods []string
	for k := range requestMethodRatio {
		methods = append(methods, k)
	}
	permanent, _ := strconv.ParseBool(os.Getenv(EnvPermanent))
	body := strings.NewReader(os.Getenv(EnvHttpRequestBody))
	slackWebHookUrl := os.Getenv(EnvSlackWebHookUrl)
	slackNotifyThreshHoldLatencyMillis, _ := strconv.Atoi(os.Getenv(EnvSlackNotifyThreshHoldLatencyMillis))
	return RuntimeInfo{
		TargetUrl:                          targetUrl,
		HttpMethods:                        methods,
		HttpHeaders:                        headers,
		ThreadNum:                          threadNum,
		TrialNum:                           trialNum,
		HttpRequestMethodRatio:             requestMethodRatio,
		Permanent:                          permanent,
		HttpRequestBody:                    body,
		SlackWebHookUrl:                    slackWebHookUrl,
		SlackNotifyThreshHoldLatencyMillis: slackNotifyThreshHoldLatencyMillis,
	}
}
