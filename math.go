package main

import (
	"fmt"
	"strconv"
) // end import

// This function accepts a list of Datum structs and computes with their values, producing the average.
func calculate_results(data []Datum) {

	var total float64 = 0.0
	var total_including_latency float64 = 0.0

	for datum := range data {
		total += data[datum].execution_time_in_seconds                    // adds up the total execution time w/o latency
		total_including_latency += data[datum].execution_time_and_latency // adds up the total execution time w/ latency
	} // end for loop

	avg_total := total / float64(len(data)) // calculates the average

	avg_total_latency := total_including_latency / float64(len(data)) // calculates the average

	fmt.Println("The average time in seconds excluding network latency is " + strconv.FormatFloat(avg_total, 'f', 10, 64))
	fmt.Println("The average time in miliseconds including network latency is " + strconv.FormatFloat(avg_total_latency, 'f', 10, 64))

} // end func
