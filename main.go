package main

import (
	"fmt"
	"os"

// "io"
)

func main() {

	fmt.Println("\tWe're inside the main func")

	data_directory_name := os.Args[len(os.Args)-1]

	data_directory, err := os.Open(data_directory_name)

	if err != nil {
		panic(err)
	} // end if

	fmt.Println("Directory loaded! Oh yeah!")

	function1(data_directory)

} // end main func
