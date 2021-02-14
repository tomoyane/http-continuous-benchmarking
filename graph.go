package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"io/ioutil"

	"github.com/wcharczuk/go-chart/v2"
)

type Graph interface {
	Output(charts []chart.Chart, startTime time.Time, endTime time.Time)

	Generate(timeRange []float64) []chart.Chart
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
			<title>Action benchmark results</title>
		</head>
		<body>
			<div>
				<p>Benchmark time: <b>%s</b> ~ <b>%s</b></p><br>`,
	startTime.String(), endTime.String()))

	for _, c := range charts {
		fmt.Println(len(charts))
		f, _ := os.Create(fmt.Sprintf("%s.png", c.Title))
		c.Render(chart.PNG, f)
		body = body + fmt.Sprintf(`<img src="./%s.png" />
        		<ul>
          			<li>
            			<b style="color: red">Red Line:</b> Maximum Percentile
					</li>
          			<li>
            			<b>Black Line:</b> Minimum Percentile
					</li>
          			<li>
            			<b style="color: orange">Orange Line:</b> Average Percentile
          			</li>
          			<li>
            			<b style="color: blue">Blue Line:</b> 99 Percentile
          			</li>
          			<li>
            			<b style="color: green">Green Line:</b> 95 Percentile
          			</li>
        		</ul>`, c.Title)
		time.Sleep(2000)
	}

	body = body + `
      		</div>
		</body>
	</html>`

	ioutil.WriteFile("index.html", []byte(body), 0644)
}

func (g GraphImpl) Generate(timeRange []float64) []chart.Chart {
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
		c := chart.Chart{
			Title: structName,
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
					XValues: timeRange,
					YValues: detail.YPercentile99,
				},
				chart.ContinuousSeries{
					Name: "95 Percentile",
					Style: chart.Style{
						StrokeColor: chart.ColorGreen,
					},
					XValues: timeRange,
					YValues: detail.YPercentile95,
				},
				chart.ContinuousSeries{
					Name: "Average Percentile",
					Style: chart.Style{
						StrokeColor: chart.ColorOrange,
					},
					XValues: timeRange,
					YValues: detail.YPercentileAvg,
				},
				chart.ContinuousSeries{
					Name: "Maximum Percentile",
					Style: chart.Style{
						StrokeColor: chart.ColorRed,
					},
					XValues: timeRange,
					YValues: detail.YPercentileMax,
				},
				chart.ContinuousSeries{
					Name: "Minimum Percentile",
					Style: chart.Style{
						StrokeColor: chart.ColorBlack,
					},
					XValues: timeRange,
					YValues: detail.YPercentileMin,
				},
			},
		}
		charts = append(charts, c)
	}
	return charts
}

func convertMetricsToGraph(metricsDetail []MetricsDetail) *GraphDetail {
	var percentile99, percentile95, percentileAvg, percentileMax, percentileMin []float64
	for _, v := range metricsDetail {
		if v.PercentileMax == 0 && v.PercentileMin == 0 && v.PercentileAvg == 0 && v.Percentile95 == 0 && v.Percentile99 == 0 {
			continue
		}
		percentile99 = append(percentile99, v.Percentile99)
		percentile95 = append(percentile95, v.Percentile95)
		percentileAvg = append(percentileAvg, v.PercentileAvg)
		percentileMax = append(percentileMax, v.PercentileMax)
		percentileMin = append(percentileMin, v.PercentileMin)
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
	}
}
