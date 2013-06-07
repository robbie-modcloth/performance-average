package main

import (
	"strconv"
	"strings"
)

type Datum struct {
	date                    string
	executionTimeInSeconds  float64
	executionTimeAndLatency float64
	errors                  string
}

// The function parses the string into individual pieces of data
// :param -- data_string, a large string with each datum delimited with "\n", and each element of the datum delimited with ","
// :return -- a slice of type datum, representing the list of all data
func parseData(dataString []string) []Datum {

	stringArray := dataString

	data := make([]Datum, len(stringArray))

	// now string_array is an array of each line, and therefore each piece of datum
	for line := range stringArray {

		parsedLine := strings.Split(stringArray[line], ",")

		if len(parsedLine) < 2 {
			continue
		} // end if

		// initializes the datum instance
		dataInstance := Datum{}
		dataInstance.date = parsedLine[0]
		dataInstance.executionTimeInSeconds, _ = strconv.ParseFloat(parsedLine[1], 64)
		dataInstance.executionTimeAndLatency, _ = strconv.ParseFloat(parsedLine[2], 64)
		dataInstance.errors = parsedLine[3]

		// now to move the instance of datum into the list of type datum
		data[line] = dataInstance
	} // end for loop

	return data

} // end function
