package main

import (
	"os"
)

// This function is the wrapper for all of the action that has to go on inside each of my mini-files.
func main() {

	dir := getDataDirectory()

	compileTheData(dir)

} // end main func

func getDataDirectory() *os.File {

	dataDirectoryName := os.Args[len(os.Args)-1]

	if dataDirectory, err := os.Open(dataDirectoryName); err != nil {
		panic(err)
	} else {
		return dataDirectory
	}

} // end func
