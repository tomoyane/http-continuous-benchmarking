package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"net/http"
)

type Calculator interface {
	GetMetricsResult() Metrics

	CalculatePerTrial(requests []int, method string, trialNum int, errData map[int]int)

	CalculateMethodErrors(srcData map[string]map[int]int, dstData map[string]map[int]int) map[string]map[int]int

	PercentileN(size int, percentile int) int
}

type Metrics struct {
	GetMetrics    []MetricsDetail
	PostMetrics   []MetricsDetail
	PutMetrics    []MetricsDetail
	PatchMetrics  []MetricsDetail
	DeleteMetrics []MetricsDetail
	TimeRange     []float64
}

// MetricsDetail Metrics per trial benchmark
// 1 trial has that below data
type MetricsDetail struct {
	Percentile99  float64
	Percentile95  float64
	PercentileAvg float64
	PercentileMax float64
	PercentileMin float64
	Rps           float64
	ErrorData     map[int]int
}

// NewCalculator Constructor
func NewCalculator(trialNum int) Calculator {
	return Metrics{
		GetMetrics:    make([]MetricsDetail, trialNum),
		PostMetrics:   make([]MetricsDetail, trialNum),
		PutMetrics:    make([]MetricsDetail, trialNum),
		PatchMetrics:  make([]MetricsDetail, trialNum),
		DeleteMetrics: make([]MetricsDetail, trialNum),
		TimeRange:     make([]float64, trialNum),
	}
}

func (m Metrics) CalculatePerTrial(requests []int, method string, trialNum int, errData map[int]int) {
	index := trialNum - 1
	samplingSize := len(requests)
	if samplingSize == 0 && len(errData) > 0 {
		detail := MetricsDetail{ErrorData: errData}
		detail.outputOnlyErrorStats(method)
		return
	}
	if samplingSize == 0 {
		return
	}

	sort.Ints(requests)
	ignore95Index := m.PercentileN(samplingSize, 95) - 1
	percentile95 := requests[ignore95Index]

	ignore99Index := m.PercentileN(samplingSize, 99) - 1
	percentile99 := requests[ignore99Index]

	var avgLatency, maxLatency, minLatency, currentRps, beforeRps int
	for i, v := range requests {
		avgLatency += v
		beforeRps = currentRps
		currentRps = v
		if i == 0 {
			minLatency = currentRps
			maxLatency = currentRps
		}
		if currentRps > beforeRps {
			maxLatency = currentRps
		}
		if currentRps < beforeRps {
			minLatency = currentRps
		}
	}

	detail := MetricsDetail{
		Percentile99:  float64(percentile99),
		Percentile95:  float64(percentile95),
		PercentileAvg: float64(avgLatency / len(requests)),
		PercentileMax: float64(maxLatency),
		PercentileMin: float64(minLatency),
		Rps:           float64(len(requests)) / float64(durationSeconds),
		ErrorData:     errData,
	}
	detail.outputStats(method)

	switch method {
	case http.MethodGet:
		m.GetMetrics[index] = detail
	case http.MethodPost:
		m.PostMetrics[index] = detail
	case http.MethodPut:
		m.PutMetrics[index] = detail
	case http.MethodPatch:
		m.PatchMetrics[index] = detail
	case http.MethodDelete:
		m.DeleteMetrics[index] = detail
	}
	m.TimeRange[index] = float64(trialNum * durationSeconds)
}

func (m Metrics) CalculateMethodErrors(srcData map[string]map[int]int, dstData map[string]map[int]int) map[string]map[int]int {
	if srcData != nil {
		for method1, errorData1 := range dstData {
			for k1, v1 := range errorData1 {
				for method2, errorData2 := range srcData {
					if method1 == method2 {
						for k2, v2 := range errorData2 {
							if k1 == k2 {
								dstData[method1][k1] = v1 + v2
							}
						}
					}
				}
			}
		}
	}
	return dstData
}

func (m Metrics) GetMetricsResult() Metrics {
	return m
}

func (m Metrics) PercentileN(size int, percentile int) int {
	n := (float64(percentile) / float64(100)) * float64(size)
	return int(math.Round(n*1) / 1)
}

func (md MetricsDetail) outputStats(method string) {
	fmt.Println(fmt.Sprintf("%s request stats information", method))
	fmt.Println(fmt.Sprintf("Latency 99  percentile: %d milliseconds", int(md.Percentile99)))
	fmt.Println(fmt.Sprintf("Latency 95  percentile: %d milliseconds", int(md.Percentile95)))
	fmt.Println(fmt.Sprintf("Latency avg percentile: %d milliseconds", int(md.PercentileAvg)))
	fmt.Println(fmt.Sprintf("Latency max percentile: %d milliseconds", int(md.PercentileMax)))
	fmt.Println(fmt.Sprintf("Latency min percentile: %d milliseconds", int(md.PercentileMin)))
	fmt.Println(fmt.Sprintf("Request per seconds:    %d", int(md.Rps)))
	if md.ErrorData != nil {
		for k, v := range md.ErrorData {
			fmt.Println(fmt.Sprintf("Error status code %d count: ", k) + strconv.Itoa(v))
		}
	}
	fmt.Println()
}

func (md MetricsDetail) outputOnlyErrorStats(method string) {
	fmt.Println(fmt.Sprintf("%s request stats information", method))
	for k, v := range md.ErrorData {
		fmt.Println(fmt.Sprintf("Error status code %d count: ", k) + strconv.Itoa(v))
	}
	fmt.Println()
}
