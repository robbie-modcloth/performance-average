package main

import (
	//"fmt"
	"os"

// "io"
)

// This function is the wrapper for all of the action that has to go on inside each of my mini-files.
func main() {

	data_directory_name := os.Args[len(os.Args)-1]
	data_directory, err := os.Open(data_directory_name)
	if err != nil {
		panic(err)
	} // end if

	// this line will accept the file object and handle almost everything else lol
	var data []Datum = capture_data(data_directory)

	calculate_results(data)

} // end main func
