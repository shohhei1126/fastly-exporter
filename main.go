package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sethvargo/go-fastly/fastly"
	fastlyExporter "github.com/shohhei1126/fastly-exporter/fastly"
	"os"
)

var (
	version string
)

var (
	listPort                 = flag.Int("port", 8080, "The port to listen on for HTTP requests.")
	metricsPath              = flag.String("path", "/metrics", "The path to return metrics.")
	useMock                  = flag.Bool("use_mock", false, "The flag to use mock of fastly client for test.")
	fastlyServiceID          = flag.String("fastly_service_id", "", "The service ID for fastly.")
	fastlyRequestIntervalSec = flag.Int("fastly_request_interval_sec", 30, "The interval to scrape for fastly real time analytics.")
	showVersion              = flag.Bool("version", false, "The flag to show version.")
)

var (
	client fastlyExporter.Client

	fastlyRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fastly_requests",
			Help: "fastly requests.",
		},
	)
)

func init() {
	prometheus.MustRegister(fastlyRequests)
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	if *useMock {
		client = fastlyExporter.NewMock()
	} else {
		if len(*fastlyServiceID) == 0 {
			panic("Fastly service ID required. set flag --fastly_service_id.")
		}
		if len(os.Getenv(fastly.APIKeyEnvVar)) == 0 {
			panic(fmt.Sprintf("Fastly API key required. Set env %v.", fastly.APIKeyEnvVar))
		}
		client = fastlyExporter.New(*fastlyServiceID)
	}

	go func() {
		for {
			fmt.Println("request to fastly...")
			res, err := client.GetLatestMetrics()
			if err != nil {
				fmt.Println("failed to get metrics of fastly" + err.Error())
			} else {
				if len(res.Data) == 1 {
					requests := float64(res.Data[0].Aggregated.Requests)
					fmt.Println(requests)
					fastlyRequests.Set(requests)
				} else {
					fmt.Println("invalid response")
				}
			}
			time.Sleep(time.Duration(*fastlyRequestIntervalSec) * time.Second)
		}
	}()

	http.Handle(*metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *listPort), nil))
}
