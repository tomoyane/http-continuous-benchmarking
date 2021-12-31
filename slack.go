package main

import (
	"bytes"
	"fmt"

	"encoding/json"
	"net/http"
)

const (
	Username = "continuous-benchmarking"
	Text = "[Alert] The threshold has been exceeded."
)

type Slack interface {
	NotifyAlert(metricsType MetricsType)
}

type SlackImpl struct {
	request *http.Request
}

// RequestBody
// Slack api request body
type RequestBody struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
}

func NewSlack(url string, channel string) Slack {
	requestBody := &RequestBody{
		Channel: channel,
		Username: Username,
		Text: Text,
	}
	jsonStr, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	return SlackImpl{
		request: req,
	}
}

func (s SlackImpl) NotifyAlert(metricsType MetricsType) {
	client := new(http.Client)
	resp, err := client.Do(s.request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(Text + fmt.Sprintf("Target metrics is %s ", metricsType.String()))
}
