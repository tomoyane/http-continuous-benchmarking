package main

import (
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	startTime := time.Now().UTC()
	metrics := Metrics{
		GetMetrics: []MetricsDetail{
			{
				PercentileMin: 1,
				PercentileMax: 30,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
			{
				PercentileMin: 4,
				PercentileMax: 21,
				PercentileAvg: 10,
				Percentile99:  16,
				Percentile95:  14,
				Rps:           100,
			},
		},
		PostMetrics: []MetricsDetail{
			{
				PercentileMin: 1,
				PercentileMax: 140,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
			{
				PercentileMin: 1,
				PercentileMax: 190,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
		},
		PutMetrics: []MetricsDetail{},
		PatchMetrics: []MetricsDetail{},
		DeleteMetrics: []MetricsDetail{
			{
				PercentileMin: 1,
				PercentileMax: 34,
				PercentileAvg: 17,
				Percentile99:  12,
				Percentile95:  13,
				Rps:           12,
			},
			{
				PercentileMin: 1,
				PercentileMax: 40,
				PercentileAvg: 19,
				Percentile99:  12,
				Percentile95:  18,
				Rps:           20,
			},
		},
		TimeRange: []float64{15, 30},
	}

	graph := NewGraph(metrics)
	charts := graph.GenerateCharts(metrics.TimeRange)

	endTime := time.Now().UTC()

	graph.Output(charts, startTime, endTime)
}

func TestGenerate_has1Chart(t *testing.T) {
	metrics := Metrics{
		GetMetrics: []MetricsDetail{
			{
				PercentileMin: 1,
				PercentileMax: 30,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
			{
				PercentileMin: 4,
				PercentileMax: 21,
				PercentileAvg: 10,
				Percentile99:  16,
				Percentile95:  14,
				Rps:           100,
			},
		},
		PostMetrics: []MetricsDetail{},
		PutMetrics: []MetricsDetail{},
		PatchMetrics: []MetricsDetail{},
		DeleteMetrics: []MetricsDetail{},
		TimeRange: []float64{15, 30},
	}

	graph := NewGraph(metrics)
	charts := graph.GenerateCharts(metrics.TimeRange)
	if len(charts) != 1 {
		t.Fatal("This test has always one chart data")
	}
}

func TestGenerate_has2Chart(t *testing.T) {
	metrics := Metrics{
		GetMetrics: []MetricsDetail{
			{
				PercentileMin: 1,
				PercentileMax: 30,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
			{
				PercentileMin: 4,
				PercentileMax: 21,
				PercentileAvg: 10,
				Percentile99:  16,
				Percentile95:  14,
				Rps:           100,
			},
		},
		PostMetrics: []MetricsDetail{
			{
				PercentileMin: 1,
				PercentileMax: 30,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
		},
		PutMetrics: []MetricsDetail{},
		PatchMetrics: []MetricsDetail{},
		DeleteMetrics: []MetricsDetail{},
		TimeRange: []float64{15, 30},
	}

	graph := NewGraph(metrics)
	charts := graph.GenerateCharts(metrics.TimeRange)

	//t.Fatal(charts[0].Series)
	if len(charts) != 2 {
		t.Fatal("This test has always two chart data")
	}
}
