package main

import (
	"os"
)

// This function is the wrapper for all of the action that has to go on inside each of my mini-files.
func main() {

	// access the name of the directory and create a file pointer to it
	data_directory_name := os.Args[len(os.Args)-1]

	if data_directory, err := os.Open(data_directory_name); err != nil {
		panic(err)
	} else {

		// this line will accept the file object and return all of the data
		var data []Datum = capture_data(data_directory)

		// this will crunch all of the numbers
		calculate_results(data)

	} // end else

} // end main func
