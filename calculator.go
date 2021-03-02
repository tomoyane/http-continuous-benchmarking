package main

import (
	"fmt"
	"math"
	"sort"

	"net/http"
)

type Calculator interface {
	GetMetricsResult() Metrics

	CalculatePerTrial(requests []int, method string, trialNum int)

	PercentileN(size int, percentile int) int
}

type Metrics struct {
	// Metrics per trial benchmark
	GetMetrics    []MetricsDetail
	PostMetrics   []MetricsDetail
	PutMetrics    []MetricsDetail
	PatchMetrics  []MetricsDetail
	DeleteMetrics []MetricsDetail
	TimeRange     []float64
}

type MetricsDetail struct {
	Percentile99  float64
	Percentile95  float64
	PercentileAvg float64
	PercentileMax float64
	PercentileMin float64
	Rps           float64
}

// Constructor
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

func (m Metrics) CalculatePerTrial(requests []int, method string, trialNum int) {
	index := trialNum - 1
	samplingSize := len(requests)
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

	fmt.Println(fmt.Sprintf("%s request stats information", method))
	fmt.Println(fmt.Sprintf("Latency 99  percentile: %d milliseconds", percentile99))
	fmt.Println(fmt.Sprintf("Latency 95  percentile: %d milliseconds", percentile95))
	fmt.Println(fmt.Sprintf("Latency avg percentile: %d milliseconds", int(float64(avgLatency / len(requests)))))
	fmt.Println(fmt.Sprintf("Latency max percentile: %d milliseconds", maxLatency))
	fmt.Println(fmt.Sprintf("Latency min percentile: %d milliseconds", minLatency))
	fmt.Println(fmt.Sprintf("Request per seconds:    %d\n", int(float64(len(requests)) / float64(durationSeconds))))

	detail := MetricsDetail{
		Percentile99:  float64(percentile99),
		Percentile95:  float64(percentile95),
		PercentileAvg: float64(avgLatency / len(requests)),
		PercentileMax: float64(maxLatency),
		PercentileMin: float64(minLatency),
		Rps:           float64(len(requests)) / float64(durationSeconds),
	}

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

func (m Metrics) GetMetricsResult() Metrics {
	return m
}

func (m Metrics) PercentileN(size int, percentile int) int {
	n := (float64(percentile) / float64(100)) * float64(size)
	return int(math.Round(n*1) / 1)
}
