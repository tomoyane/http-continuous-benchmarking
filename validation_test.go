package main

import (
	"os"
	"testing"
)

func TestValidateEnv(t *testing.T) {
	os.Setenv(EnvSlackNotifyThreshHoldLatencyMillis, "invalid")
	os.Setenv(EnvTargetUrl, "")
	os.Setenv(EnvHttpHeaders, "")
	os.Setenv(EnvReqHttpMethodRatio, "")
	os.Setenv(EnvHttpRequestBody, "")
	os.Setenv(EnvThreadNum, "")
	os.Setenv(EnvTrialNum, "")
	result := ValidateEnv()
	if result == nil {
		t.Fatalf("Expect required validation error.")
	}
}

func TestValidateTargetUrl(t *testing.T) {
	os.Setenv(EnvTargetUrl, "")
	result := validateTargetUrl()
	if result == nil {
		t.Fatalf("Expect empty validation error. %s", EnvTargetUrl)
	}

	os.Setenv(EnvTargetUrl, "invalid_url")
	result = validateTargetUrl()
	if result == nil {
		t.Fatalf("Expect prefix validation error. %s ", EnvTargetUrl)
	}
}

func TestValidateHttpHeader(t *testing.T) {
	os.Setenv(EnvHttpHeaders, "")
	result := validateHttpHeaders()
	if result == nil {
		t.Fatalf("Expect empty validation error. %s", EnvHttpHeaders)
	}

	os.Setenv(EnvHttpHeaders, "Not map data")
	result = validateHttpHeaders()
	if result == nil {
		t.Fatalf("Expect map parser validation error. %s", EnvHttpHeaders)
	}
}

func TestValidateReqHttpMethodPercentage(t *testing.T) {
	os.Setenv(EnvReqHttpMethodRatio, "")
	result := validateReqHttpMethodRatio()
	if result == nil {
		t.Fatalf("Expect empty validation error. %s", EnvReqHttpMethodRatio)
	}

	os.Setenv(EnvReqHttpMethodRatio, "Not map data")
	result = validateReqHttpMethodRatio()
	if result == nil {
		t.Fatalf("Expect map parser validation error. %s", EnvReqHttpMethodRatio)
	}

	os.Setenv(EnvReqHttpMethodRatio, `{"GET": 7, "PUT":2, "POST":2}`)
	result = validateReqHttpMethodRatio()
	if result == nil {
		t.Fatalf("Expect not filled percentage validation error. Just 10 ratio. %s", EnvReqHttpMethodRatio)
	}
}

func TestValidateHttpRequestBody(t *testing.T) {
	os.Setenv(EnvHttpRequestBody, "")
	result := validateHttpRequestBody()
	if result == nil {
		t.Fatalf("Expect empty validation error. %s", EnvHttpRequestBody)
	}

	os.Setenv(EnvHttpRequestBody, "Not map data")
	result = validateHttpRequestBody()
	if result == nil {
		t.Fatalf("Expect map parser validation error. %s", EnvHttpRequestBody)
	}
}

func TestValidateThreadNum(t *testing.T) {
	os.Setenv(EnvThreadNum, "")
	result := validateThreadNum()
	if result == nil {
		t.Fatalf("Expect empty validation error. %s", EnvThreadNum)
	}

	os.Setenv(EnvThreadNum, "string")
	result = validateThreadNum()
	if result == nil {
		t.Fatalf("Expect number validation error. %s", EnvThreadNum)
	}
}

func TestValidateTrialNum(t *testing.T) {
	os.Setenv(EnvTrialNum, "")
	result := validateTrialNum()
	if result == nil {
		t.Fatalf("Expect empty validation error. %s", EnvTrialNum)
	}

	os.Setenv(EnvTrialNum, "string")
	result = validateTrialNum()
	if result == nil {
		t.Fatalf("Expect number validation error. %s", EnvTrialNum)
	}
}

func TestValidatePermanent(t *testing.T) {
	os.Setenv(EnvPermanent, "invalid")
	result := validatePermanent()
	if result == nil {
		t.Fatalf("Expect invalid validation error. %s", EnvPermanent)
	}

	os.Setenv(EnvPermanent, "invalid")
	result = validatePermanent()
	if result == nil {
		t.Fatalf("Expect allowed true or false validation error. %s", EnvPermanent)
	}
}

func TestValidateSlackNotifyThreshHoldLatencyMillis(t *testing.T) {
	os.Setenv(EnvSlackNotifyThreshHoldLatencyMillis, "invalid")
	result := validateAlert()
	if result == nil {
		t.Fatalf("Expect invalid validation error. %s", EnvSlackNotifyThreshHoldLatencyMillis)
	}

	os.Setenv(EnvSlackNotifyThreshHoldLatencyMillis, "string")
	result = validateAlert()
	if result == nil {
		t.Fatalf("Expect number validation error. %s", EnvSlackNotifyThreshHoldLatencyMillis)
	}
}
