package main

import (
	"fmt"
	"strconv"
) // end import

func calculate_results(data []Datum) {

	var total float64 = 0.0
	var total_including_latency float64 = 0.0
	for datum := range data {
		total += data[datum].execution_time_in_seconds
		total_including_latency += data[datum].execution_time_and_latency
	} // end for loop

	avg_total := total / float64(len(data))

	avg_total_latency := total_including_latency / float64(len(data))

	fmt.Println("The average time in seconds excluding network latency is " + strconv.FormatFloat(avg_total, 'f', 10, 64))
	fmt.Println("The average time in miliseconds including network latency is " + strconv.FormatFloat(avg_total_latency, 'f', 10, 64))

} // end func
