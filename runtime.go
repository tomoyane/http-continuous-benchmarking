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
	EnableAlert                        bool
	SlackWebHookUrl                    string
	SlackChannel                       string
	SlackNotifyThreshHoldLatencyMillis int
	SlackNotifyThreshHoldRps           int
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

	isEnable, _ := strconv.ParseBool(os.Getenv(EnvEnableAlert))
	slackWebHookUrl := os.Getenv(EnvSlackWebHookUrl)
	slackChannel := os.Getenv(EnvSlackChannel)
	if !strings.HasPrefix(slackChannel, "#") {
		slackChannel = "#" + slackChannel
	}
	slackNotifyThreshHoldLatencyMillis, _ := strconv.Atoi(os.Getenv(EnvSlackNotifyThreshHoldLatencyMillis))
	slackNotifyThreshHoldRps, _ := strconv.Atoi(os.Getenv(EnvSlackNotifyThreshHoldRps))
	return RuntimeInfo{
		TargetUrl:                          targetUrl,
		HttpMethods:                        methods,
		HttpHeaders:                        headers,
		ThreadNum:                          threadNum,
		TrialNum:                           trialNum,
		HttpRequestMethodRatio:             requestMethodRatio,
		Permanent:                          permanent,
		HttpRequestBody:                    body,
		EnableAlert:                        isEnable,
		SlackWebHookUrl:                    slackWebHookUrl,
		SlackChannel:                       slackChannel,
		SlackNotifyThreshHoldLatencyMillis: slackNotifyThreshHoldLatencyMillis,
		SlackNotifyThreshHoldRps:           slackNotifyThreshHoldRps,
	}
}
