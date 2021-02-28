package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"io/ioutil"
	"net/http"

	"github.com/wcharczuk/go-chart/v2"
)

type Graph interface {
	Output(charts []chart.Chart, startTime time.Time, endTime time.Time)

	GenerateCharts(timeRange []float64) []chart.Chart
}

type GraphImpl struct {
	GetGraphDetail    *GraphDetail
	PostGraphDetail   *GraphDetail
	PutGraphDetail    *GraphDetail
	PatchGraphDetail  *GraphDetail
	DeleteGraphDetail *GraphDetail
}

type GraphDetail struct {
	YPercentile99  []float64
	YPercentile95  []float64
	YPercentileAvg []float64
	YPercentileMax []float64
	YPercentileMin []float64
	Rps            []float64
}

func NewGraph(metrics Metrics) Graph {
	getGraphDetail := convertMetricsToGraph(metrics.GetMetrics)
	postGraphDetail := convertMetricsToGraph(metrics.PostMetrics)
	putGraphDetail := convertMetricsToGraph(metrics.PutMetrics)
	patchGraphDetail := convertMetricsToGraph(metrics.PatchMetrics)
	deleteGraphDetail := convertMetricsToGraph(metrics.DeleteMetrics)
	return GraphImpl{
		GetGraphDetail:    getGraphDetail,
		PostGraphDetail:   postGraphDetail,
		PutGraphDetail:    putGraphDetail,
		PatchGraphDetail:  patchGraphDetail,
		DeleteGraphDetail: deleteGraphDetail,
	}
}

func (g GraphImpl) Output(charts []chart.Chart, startTime time.Time, endTime time.Time) {
	body := fmt.Sprintf(fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>HTTP CB results</title>
	<style type="text/css">
		.title {text-align: center}
		.container {display: flex; margin-left:100px; margin-right:100px}
		.child {flex-grow: 1;}
		.graph-img {width: 600px}
	</style>
	</head>
	<body>
	<div>
	<p class="title">Benchmark time: <b>%s</b> ~ <b>%s</b></p><br>`,
		startTime.String(), endTime.String()))

	for i := 0; i < len(charts)/2; i++ {
		index := i * 2

		body = body + `<div class="container">`
		f, _ := os.Create(fmt.Sprintf("%s.png", charts[index].Title))
		charts[index].Render(chart.PNG, f)
		body = body + fmt.Sprintf(`
		<div class="child">
		<img src="./%s.png" class="graph-img" />
		<ul>
		<li><b style="color: red">Red Line:</b> Maximum Percentile</li>
		<li><b>Black Line:</b> Minimum Percentile</li>
		<li><b style="color: orange">Orange Line:</b> Average Percentile</li>
		<li><b style="color: blue">Blue Line:</b> 99 Percentile</li>
		<li><b style="color: green">Green Line:</b> 95 Percentile</li>
		</ul></div>`, charts[index].Title)

		f, _ = os.Create(fmt.Sprintf("%s.png", charts[index+1].Title))
		charts[index+1].Render(chart.PNG, f)
		body = body + fmt.Sprintf(`
		<div class="child">
		<img src="./%s.png" class="graph-img" />
		<ul>
		<li><b>Black Line:</b> Request per seconds</li>
		</ul></div>`, charts[index+1].Title)
		body = body + `</div>`
	}

	body = body + `</div></body></html>`

	ioutil.WriteFile("index.html", []byte(body), 0644)
}

func (g GraphImpl) GenerateCharts(timeRange []float64) []chart.Chart {
	t := reflect.TypeOf(g)
	elem := reflect.ValueOf(&g).Elem()
	cnt := elem.NumField()
	var charts []chart.Chart

	for i := 0; i < cnt; i++ {
		structName := t.Field(i).Name
		structData := elem.Field(i)
		detail := structData.Interface().(*GraphDetail)
		if detail == nil {
			continue
		}
		charts = append(charts, latencyChart(structName, timeRange, *detail))
		charts = append(charts, rpsChart(structName, timeRange, *detail))
	}
	return charts
}

func convertMetricsToGraph(metricsDetail []MetricsDetail) *GraphDetail {
	var percentile99, percentile95, percentileAvg, percentileMax, percentileMin, rps []float64
	for _, v := range metricsDetail {
		if v.PercentileMax == 0 && v.PercentileMin == 0 && v.PercentileAvg == 0 && v.Percentile95 == 0 && v.Percentile99 == 0 {
			continue
		}
		percentile99 = append(percentile99, v.Percentile99)
		percentile95 = append(percentile95, v.Percentile95)
		percentileAvg = append(percentileAvg, v.PercentileAvg)
		percentileMax = append(percentileMax, v.PercentileMax)
		percentileMin = append(percentileMin, v.PercentileMin)
		rps = append(rps, v.Rps)
	}

	if len(percentileMax) == 0 || len(percentileMin) == 0 || len(percentileAvg) == 0 || len(percentile99) == 0 || len(percentile95) == 0 {
		return nil
	}

	return &GraphDetail{
		YPercentile95:  percentile95,
		YPercentile99:  percentile99,
		YPercentileAvg: percentileAvg,
		YPercentileMax: percentileMax,
		YPercentileMin: percentileMin,
		Rps:            rps,
	}
}

func getChartPrefixTitleFrom(structName string) string {
	var method string
	if strings.Contains(strings.ToLower(structName), strings.ToLower(http.MethodGet)) {
		method = http.MethodGet
	} else if strings.Contains(strings.ToLower(structName), strings.ToLower(http.MethodPost)) {
		method = http.MethodPost
	} else if strings.Contains(strings.ToLower(structName), strings.ToLower(http.MethodPut)) {
		method = http.MethodPut
	} else if strings.Contains(strings.ToLower(structName), strings.ToLower(http.MethodPatch)) {
		method = http.MethodPatch
	} else if strings.Contains(strings.ToLower(structName), strings.ToLower(http.MethodDelete)) {
		method = http.MethodDelete
	}
	return method
}

func latencyChart(structName string, x []float64, detail GraphDetail) chart.Chart {
	c := chart.Chart{
		Title: getChartPrefixTitleFrom(structName) + " Latency",
		XAxis: chart.XAxis{
			Name: "Time (sec)",
		},
		YAxis: chart.YAxis{
			Name: "Latency (ms)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "99 Percentile",
				Style: chart.Style{
					StrokeColor: chart.ColorBlue,
				},
				XValues: x,
				YValues: detail.YPercentile99,
			},
			chart.ContinuousSeries{
				Name: "95 Percentile",
				Style: chart.Style{
					StrokeColor: chart.ColorGreen,
				},
				XValues: x,
				YValues: detail.YPercentile95,
			},
			chart.ContinuousSeries{
				Name: "Average Percentile",
				Style: chart.Style{
					StrokeColor: chart.ColorOrange,
				},
				XValues: x,
				YValues: detail.YPercentileAvg,
			},
			chart.ContinuousSeries{
				Name: "Maximum Percentile",
				Style: chart.Style{
					StrokeColor: chart.ColorRed,
				},
				XValues: x,
				YValues: detail.YPercentileMax,
			},
			chart.ContinuousSeries{
				Name: "Minimum Percentile",
				Style: chart.Style{
					StrokeColor: chart.ColorBlack,
				},
				XValues: x,
				YValues: detail.YPercentileMin,
			},
		},
	}
	return c
}

func rpsChart(structName string, x []float64, detail GraphDetail) chart.Chart {
	c := chart.Chart{
		Title: getChartPrefixTitleFrom(structName) + " Request per seconds",
		XAxis: chart.XAxis{
			Name: "Time (sec)",
		},
		YAxis: chart.YAxis{
			Name: "Request per seconds",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "Request per seconds",
				Style: chart.Style{
					StrokeColor: chart.ColorBlack,
				},
				XValues: x,
				YValues: detail.Rps,
			},
		},
	}
	return c
}
