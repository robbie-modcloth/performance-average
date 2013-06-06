package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// This function consolidates the data in all of the individual files into a buffer, then into a string
// After, the data is converted from type string into type Datum, which is a struct representing the four fields of the CSV
func capture_data(data_directory *os.File) []Datum {
	data_string := compile_the_data(data_directory)
	return parse_data(data_string)
}

// This accepts a directory from the main function
// The directory has all of the data stored in the series of files in it.
func compile_the_data(data_directory *os.File) string {

	// initializes a buffer to store all the data
	var buf_arr []byte
	buf := bytes.NewBuffer(buf_arr)

	// this is to get the pwd from within the new directory
	base_dir := data_directory.Name()

	// this reads all of the file names in the directory and stores them in a slice
	if partial_file_names, err := data_directory.Readdirnames(-1); err != nil {
		panic(err)

	} else {

		list_of_file_names := get_complete_file_names(base_dir, partial_file_names)

		// for each file in the directory...
		for file := range list_of_file_names {

			// open the file
			my_file, err := os.Open(list_of_file_names[file])
			if err != nil {
				panic(err)
			} // end if
			defer my_file.Close()

			// then grab all the data
			buf.Write(pull_data(my_file))

		} // end for loop

	} // end else

	fmt.Println("The data has been successfully compiled!")
	return buf.String()

} // end function

// This function accepts one file and reads the whole thing into a byte array and returns it.
func pull_data(file *os.File) []byte {

	byte_array, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		panic("IO ERROR")
	} // end if

	return byte_array

} // end function

func get_complete_file_names(base_dir string, partial_file_names []string) []string {

	names := make([]string, len(partial_file_names))

	for index := range partial_file_names {
		names[index] = base_dir + "/" + partial_file_names[index]
	} // end for loop

	return names

} // end func
