/*
A simple sample of how an Prometheus should be implemented in Go
For more information, refer to https://pkg.go.dev/github.com/prometheus/client_golang/prometheus
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var seed = 10
var disks = []string{"sda", "sdb", "sdc"}
var diskDesc = map[string]*prometheus.Desc{}

func init() {
	rand.Seed(int64(seed))
}

func descInit(disks []string) {
	for _, disk := range disks {
		diskDesc[disk] = prometheus.NewDesc(disk, disk+" metrics", []string{"name", "device"}, nil)
	}
}

type diskCollector struct {
	disks []string
}

func (dc *diskCollector) Describe(ch chan<- *prometheus.Desc) {
	descInit(dc.disks)

	for _, desc := range diskDesc {
		ch <- desc
	}
}

func (dc *diskCollector) Collect(ch chan<- prometheus.Metric) {
	for disk, desc := range diskDesc {
		metric := rand.Int63n(10000)
		ch <- prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(metric),
			disk,
			fmt.Sprintf("/dev/%s1", disk),
		)
	}
}

func main() {
	dc := diskCollector{disks: disks}
	reg := prometheus.NewRegistry()
	reg.MustRegister(&dc)

	// Add process and go internal stats information
	reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}), prometheus.NewGoCollector())

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
