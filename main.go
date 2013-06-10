package main

import (
	"os"
)

// This function is the wrapper for all of the action that has to go on inside each of my mini-files.
func main() {

	// access the name of the directory and create a file pointer to it
	dataDirectoryName := os.Args[len(os.Args)-1]

	if dataDirectory, err := os.Open(dataDirectoryName); err != nil {
		panic(err)
	} else {

		compileTheData(dataDirectory)

	} // end else

} // end main func
