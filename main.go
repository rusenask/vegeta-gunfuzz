package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)



func main() {
	rate := uint64(100) // per second
	duration := 4 * time.Second
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
