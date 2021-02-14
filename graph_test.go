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
				PercentileMax: 30,
				PercentileAvg: 10,
				Percentile99:  19,
				Percentile95:  18,
				Rps:           20,
			},
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
	charts := graph.Generate(metrics.TimeRange)

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
	charts := graph.Generate(metrics.TimeRange)
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
	charts := graph.Generate(metrics.TimeRange)

	//t.Fatal(charts[0].Series)
	if len(charts) != 2 {
		t.Fatal("This test has always two chart data")
	}
}
