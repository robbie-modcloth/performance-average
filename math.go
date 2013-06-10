package main

import (
	"fmt"
	"strconv"
) // end import

// This function accepts a list of Datum structs and computes with their values, producing the average.
func calculateAverageDuration(data *chanList) {

	var total float64 = 0.0
	var totalIncludingLatency float64 = 0.0
	var totalNumData int = 0

	fmt.Println(len(data.resultsTotal1))

	for datum := 0; datum < len(data.resultsTotal1); datum++ {
		total += <-data.resultsTotal1                 // adds up the total execution time w/o latency
		totalIncludingLatency += <-data.resultsTotal2 // adds up the total execution time w/ latency
		totalNumData += <-data.resultsLen
		///fmt.Println(strconv.FormatFloat(total, 'f', 'f', 64))
	} // end for loop

	avgTotal := total / float64(totalNumData) // calculates the average

	avgTotalLatency := totalIncludingLatency / float64(totalNumData) // calculates the average

	fmt.Println("The average time in seconds excluding network latency is " + strconv.FormatFloat(avgTotal, 'f', 10, 64))
	fmt.Println("The average time in miliseconds including network latency is " + strconv.FormatFloat(avgTotalLatency, 'f', 10, 64))

} // end func
