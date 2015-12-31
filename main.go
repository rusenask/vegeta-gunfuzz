package main

import (
	"fmt"
	"time"
	"flag"

	vegeta "github.com/tsenart/vegeta/lib"
)

type AppConfig struct {
	totalTime int
	ratePerSecond int
}

var appConfig AppConfig


func main() {
	// parsing flags
	var totalTime = flag.Int("time", 4, "time to run tests")
	var ratePerSecond = flag.Int("rate", 100, "time to run tests")
	flag.Parse() // parse the flag

	appConfig.totalTime = *totalTime
	appConfig.ratePerSecond = *ratePerSecond

	rate := uint64(appConfig.ratePerSecond) // per second
	duration := time.Duration(appConfig.totalTime) * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://unfuzzservice2.azurewebsites.net/api/categories/filter?q=led",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
