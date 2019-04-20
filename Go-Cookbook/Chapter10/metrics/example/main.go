package main

import (
	"fmt"
	"net/http"

	"../../metrics"
)

func main() {
	// handler to populate metrics
	http.HandleFunc("/counter", metrics.CounterHandler)
	http.HandleFunc("/timer", metrics.TimerHandler)
	http.HandleFunc("/report", metrics.ReportHandler)
	fmt.Println("listening on :8080")
	panic(http.ListenAndServe(":8080", nil))
}
