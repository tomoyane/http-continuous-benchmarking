package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"encoding/json"
	"net/http"
)

var allowedHttpMethod = []string{
	http.MethodGet,
	http.MethodPatch,
	http.MethodPut,
	http.MethodPost,
	http.MethodDelete,
}

func ValidateEnv() map[string]string {
	result := make(map[string]string)
	validateTargetUrl(result)
	validateHttpMethods(result)
	validateHttpHeaders(result)
	validateReqHttpMethodPercentage(result)
	validateHttpRequestBody(result)
	validateThreadNum(result)
	validateLoadTimeSeconds(result)
	validatePermanent(result)
	validateSlackNotifyThreshHoldLatencyMillis(result)
	return result
}

// Validate TARGET_URL env
func validateTargetUrl(result map[string]string) map[string]string {
	env := os.Getenv(EnvTargetUrl)
	if validateEmpty(env) {
		result[EnvTargetUrl] = fmt.Sprintf("Environment valiable %s is required.", EnvTargetUrl)
		return result
	}
	if !strings.HasPrefix(env, "http") || !strings.HasPrefix(env, "https") {
		result[EnvTargetUrl] = fmt.Sprintf("Environment valiable %s has only http or https protocol.", EnvTargetUrl)
		return result
	}
	return nil
}

// Validate HTTP_METHODS env
func validateHttpMethods(result map[string]string) map[string]string {
	env := os.Getenv(EnvHttpMethods)
	if validateEmpty(env) {
		result[EnvHttpMethods] = fmt.Sprintf("Environment valiable %s is required.", EnvHttpMethods)
		return result
	}
	isContain := false
	methods := strings.Split(env, ",")
	for _, v := range allowedHttpMethod {
		for _, m := range methods {
			if strings.EqualFold(v, m) {
				isContain = true
				break
			}
		}
	}
	if !isContain {
		result[EnvHttpMethods] = fmt.Sprintf("Environment valiable %s is only supprt %v.", EnvHttpMethods, allowedHttpMethod)
		return result
	}
	return nil
}

// Validate HTTP_HEADERS env
func validateHttpHeaders(result map[string]string) map[string]string {
	env := os.Getenv(EnvHttpHeaders)
	if validateEmpty(env) {
		result[EnvHttpHeaders] = fmt.Sprintf("Environment valiable %s is required.", EnvHttpHeaders)
		return result
	}
	headers := make(map[string]interface{})
	if err := json.Unmarshal([]byte(env), &headers); err != nil {
		result[EnvHttpHeaders] = fmt.Sprintf("Environment valiable %s not hashmap structure.", EnvHttpHeaders)
		return result
	}
	return nil
}

// Validate REQ_HTTP_METHOD_PERCENTAGES env
func validateReqHttpMethodPercentage(result map[string]string) map[string]string {
	methods := strings.Split(os.Getenv(EnvHttpMethods), ",")
	if len(methods) > 1 {
		env := os.Getenv(EnvReqHttpMethodPercentages)
		if validateEmpty(env) {
			result[EnvReqHttpMethodPercentages] = fmt.Sprintf("Environment valiable %s is required because method is multiple.", EnvReqHttpMethodPercentages)
			return result
		}
		percentages := make(map[string]interface{})
		if err := json.Unmarshal([]byte(env), &percentages); err != nil {
			result[EnvReqHttpMethodPercentages] = fmt.Sprintf("Environment valiable %s not hashmap structure.", EnvReqHttpMethodPercentages)
			return result
		}
	}
	return nil
}

// Validate HTTP_REQ_BODY env
func validateHttpRequestBody(result map[string]string) map[string]string {
	env := os.Getenv(EnvHttpRequestBody)
	if validateEmpty(env) {
		result[EnvHttpRequestBody] = fmt.Sprintf("Environment valiable %s is required.", EnvHttpRequestBody)
		return result
	}
	body := make(map[string]interface{})
	if err := json.Unmarshal([]byte(env), &body); err != nil {
		result[EnvHttpRequestBody] = fmt.Sprintf("Environment valiable %s not hashmap structure.", EnvHttpRequestBody)
		return result
	}
	return nil
}

// Validate THREAD_NUM env
func validateThreadNum(result map[string]string) map[string]string {
	env := os.Getenv(EnvThreadNum)
	if validateEmpty(env) {
		result[EnvThreadNum] = fmt.Sprintf("Environment valiable %s is required.", EnvThreadNum)
		return result
	}
	if _, err := strconv.Atoi(env); err != nil {
		result[EnvThreadNum] = fmt.Sprintf("Environment valiable %s is not number.", EnvThreadNum)
		return result
	}
	return nil
}

// Validate LOAD_TIME_SECONDS env
func validateLoadTimeSeconds(result map[string]string) map[string]string {
	env := os.Getenv(EnvTrialNum)
	if validateEmpty(env) {
		result[EnvTrialNum] = fmt.Sprintf("Environment valiable %s is required.", EnvTrialNum)
		return result
	}
	if _, err := strconv.Atoi(env); err != nil {
		result[EnvTrialNum] = fmt.Sprintf("Environment valiable %s is not number.", EnvTrialNum)
		return result
	}
	return nil
}

// Validate PERMANENT env
func validatePermanent(result map[string]string) map[string]string {
	env := os.Getenv(EnvPermanent)
	if validateEmpty(env) {
		return nil
	}
	if !strings.EqualFold(env, "true") || !strings.EqualFold(env, "false") {
		result[EnvPermanent] = fmt.Sprintf("Environment valiable %s is true or false.", EnvPermanent)
		return result
	}
	return nil
}

// Validate SLACK_NOTIFY_THRESHOLD_LATENCY_MILLIS env
func validateSlackNotifyThreshHoldLatencyMillis(result map[string]string) map[string]string {
	env := os.Getenv(EnvSlackNotifyThreshHoldLatencyMillis)
	if validateEmpty(env) {
		return nil
	}
	if _, err := strconv.Atoi(env); err != nil {
		result[EnvSlackNotifyThreshHoldLatencyMillis] = fmt.Sprintf("Environment valiable %s is not number.", EnvSlackNotifyThreshHoldLatencyMillis)
		return result
	}
	return nil
}

// Validate empty
func validateEmpty(data string) bool {
	if data == "" {
		return true
	}
	return false
}
