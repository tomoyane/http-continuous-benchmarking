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
	HttpRequestMethodPercentage        map[string]int
	Permanent                          bool
	HttpRequestBody                    io.Reader
	SlackWebHookUrl                    string
	SlackNotifyThreshHoldLatencyMillis int
}

// New RuntimeInfo from environment variable
func NewRuntimeInfo() RuntimeInfo {
	targetUrl := os.Getenv(EnvTargetUrl)
	methods := strings.Split(os.Getenv(EnvHttpMethods), ",")
	headers := make(map[string]string)
	json.Unmarshal([]byte(os.Getenv(EnvHttpHeaders)), &headers)
	threadNum, _ := strconv.Atoi(os.Getenv(EnvThreadNum))
	trialNum, _ := strconv.Atoi(os.Getenv(EnvTrialNum))
	percentages := make(map[string]int)
	json.Unmarshal([]byte(os.Getenv(EnvReqHttpMethodPercentages)), &percentages)
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
		HttpRequestMethodPercentage:        percentages,
		Permanent:                          permanent,
		HttpRequestBody:                    body,
		SlackWebHookUrl:                    slackWebHookUrl,
		SlackNotifyThreshHoldLatencyMillis: slackNotifyThreshHoldLatencyMillis,
	}
}
