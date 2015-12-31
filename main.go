package main

import (
	"fmt"
	"time"
	"flag"

	vegeta "github.com/tsenart/vegeta/lib"
)



func main() {
	// parsing flags
	var totalTime = flag.Int("time", 4, "time to run tests")
	flag.Parse() // parse the flag

	rate := uint64(20) // per second
	duration := *totalTime * time.Second
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
