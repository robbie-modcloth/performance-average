package main

import (
	"fmt"
	"os"
	"bufio"
)

// This accepts a directory from the main function
// The directory has all of the data stored in the series of files in it.
func compile_the_data(data_directory *os.File) {

  // this reads all of the file names in the directory and stores them in a slice
	file_names, err := (*data_directory).Readdirnames(-1)
	if err != nil {
		panic(err)
		fmt.Println("PAAAAAANNNNIIIIIIC!")
	} // end if

  // this just prints all of the file names to std out
	fmt.Println(file_names)

  // this is to get the pwd from within the new directory
	base_dir := (*data_directory).Name()

  // initialized a string which will store all of the data as one big string
  // TODO: REPLACE WITH A BUFFER
	data_string := ""

  // for each file in the directory...
	for file := range file_names {

    // get its full name, composed of the name of the directory above concat the name of that file from within the directory
		individual_file := base_dir + file_names[file]
		my_file, _ := os.Open(individual_file)
		
    
    // TODO: fix this!
    // defer my_file.close()
    // this should be a buffer...
		data_string += pull_data(my_file)

	} // end for loop

} // end function

// This function accepts one file and reads each line into a buffer, then returns the buffer.... errr... string.
func pull_data(file *os.File) string {

	reader := bufio.NewReader(file)

  // TODO fix the next two lines
	fmt.Println(reader.Buffered())
	array := make([5]byte)

	return ""

} // end function
