package main

import (
	"testing"

	"net/http"
)

var testGetRequestLatencies = []int{
	10, 20, 24, 30, 12, 27, 18, 57, 9, 14,
	12, 10, 54, 33, 23, 54, 3, 4, 5, 7, 1,
	10, 15, 12, 34, 21, 12, 34, 3, 4, 6, 7,
}

func TestCalculate(t *testing.T) {
	calculator := NewCalculator(15)
	calculator.CalculatePerTrial(testGetRequestLatencies, http.MethodGet, 1)
	metrics := calculator.GetMetricsResult()

	if int(metrics.GetMetrics[0].Rps) != (len(testGetRequestLatencies) / durationSeconds) {
		t.Fatalf("Not equals RPS. Expect = %d, Actual = %d", len(testGetRequestLatencies), int(metrics.GetMetrics[0].Rps))
	}

	if int(metrics.GetMetrics[0].PercentileAvg) != 18 {
		t.Fatalf("Not equals avg percentile. Expect = 18, Actual = %d", int(metrics.GetMetrics[0].PercentileAvg))
	}

	if int(metrics.GetMetrics[0].PercentileMin) != 1 {
		t.Fatalf("Not equals min percentile. Expect = 1, Actual = %d", int(metrics.GetMetrics[0].PercentileMin))
	}

	if int(metrics.GetMetrics[0].PercentileMax) != 57 {
		t.Fatalf("Not equals max percentile. Expect = 57, Actual = %d", int(metrics.GetMetrics[0].PercentileMax))
	}

	if int(metrics.GetMetrics[0].Percentile95) != 54 {
		t.Fatalf("Not equals 95 percentile. Expect = %d, Actual = %d", 54, int(metrics.GetMetrics[0].Percentile95))
	}

	if int(metrics.GetMetrics[0].Percentile99) != 57 {
		t.Fatalf("Not equals 99 percentile. Expect = %d, Actual = %d", 57, int(metrics.GetMetrics[0].Percentile99))
	}

	if int(metrics.GetMetrics[1].Percentile95) != 0 && int(metrics.GetMetrics[1].Percentile99) != 0 {
		t.Fatalf("Cap is always zero when index is not 0")
	}
}

func TestPercentileN(t *testing.T) {
	calculator := NewCalculator(1)
	ignoreIndex := calculator.PercentileN(8, 95)
	if ignoreIndex != 8 {
		t.Fatalf("Failed to percentile calculator.")
	}
}
