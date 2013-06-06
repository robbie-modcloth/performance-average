package main

import (
	"strconv"
	"strings"
)

type Datum struct {
	date                       string
	execution_time_in_seconds  float64
	execution_time_and_latency float64
	errors                     string
}

// The function parses the string into individual pieces of data
// :param -- data_string, a large string with each datum delimited with "\n", and each element of the datum delimited with ","
// :return -- a slice of type datum, representing the list of all data
func parse_data(data_string string) []Datum {

	string_array := strings.Split(data_string, "\n")

	data := make([]Datum, len(string_array))

	// now string_array is an array of each line, and therefore each piece of datum
	for line := range string_array {

		parsed_line := strings.Split(string_array[line], ",")

		if len(parsed_line) < 2 {
			continue
		} // end if

		// initializes the datum instance
		data_instance := Datum{}
		data_instance.date = parsed_line[0]
		data_instance.execution_time_in_seconds, _ = strconv.ParseFloat(parsed_line[1], 64)
		data_instance.execution_time_and_latency, _ = strconv.ParseFloat(parsed_line[2], 64)
		data_instance.errors = parsed_line[3]

		// now to move the instance of datum into the list of type datum
		data[line] = data_instance
	} // end for loop

	return data

} // end function
