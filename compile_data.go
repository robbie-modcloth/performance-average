package main

import (
	//	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

//	"strings"
)

func capture_data(data_directory *os.File) []Datum {
	data_string := compile_the_data(data_directory)
	return parse_data(data_string)
}

// This accepts a directory from the main function
// The directory has all of the data stored in the series of files in it.
func compile_the_data(data_directory *os.File) string {

	// this reads all of the file names in the directory and stores them in a slice
	file_names, err := data_directory.Readdirnames(-1)
	if err != nil {
		panic(err)
		fmt.Println("PAAAAAANNNNIIIIIIC!")
	} // end if

	// this is to get the pwd from within the new directory
	base_dir := (*data_directory).Name()

	// initializes a buffer to store all the data
	var buf_arr []byte
	buf := bytes.NewBuffer(buf_arr)

	// for each file in the directory...
	for file := range file_names {

		// get its full name, composed of the name of the directory above concat the name of that file from within the directory
		individual_file := base_dir + "/" + file_names[file]
		my_file, err := os.Open(individual_file)
		if err != nil {
			panic(err)
		} // end if

		defer my_file.Close()
		// this should be a buffer...
		buf.Write(pull_data(my_file))

	} // end for loop

	fmt.Println("The buffer has successfully been created!")

	return buf.String()

} // end function

// This function accepts one file and reads each line into a buffer, then returns the buffer.... errr... string.
func pull_data(file *os.File) []byte {

	byte_array, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		panic("IO ERROR")
	} // end if

	return byte_array

} // end function
