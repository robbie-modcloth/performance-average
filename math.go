package main

import (
	"fmt"
	"strconv"
) // end import

// This function accepts a list of Datum structs and computes with their values, producing the average.
func calculateAverageDuration(data []Datum) {

	var total float64 = 0.0
	var totalIncludingLatency float64 = 0.0

	for datum := range data {
		total += data[datum].executionTimeInSeconds                  // adds up the total execution time w/o latency
		totalIncludingLatency += data[datum].executionTimeAndLatency // adds up the total execution time w/ latency
	} // end for loop

	avgTotal := total / float64(len(data)) // calculates the average

	avgTotalLatency := totalIncludingLatency / float64(len(data)) // calculates the average

	fmt.Println("The average time in seconds excluding network latency is " + strconv.FormatFloat(avgTotal, 'f', 10, 64))
	fmt.Println("The average time in miliseconds including network latency is " + strconv.FormatFloat(avgTotalLatency, 'f', 10, 64))

} // end func
