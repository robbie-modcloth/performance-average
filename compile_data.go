package main

import (
	"bufio"
	"fmt"
	"os"
)

func captureData(dataDirectory *os.File) []Datum {
	dataSlice := compileTheData(dataDirectory)
	return parseData(dataSlice)
}

func compileTheData(dataDirectory *os.File) []string {

	var dataLines []string

	if partialFileNames, err := dataDirectory.Readdirnames(-1); err != nil {
		panic(err)
	} else {

		baseDir := dataDirectory.Name()
		listOfFileNames := getCompleteFileNames(baseDir, partialFileNames)
		dataLines = make([]string, len(listOfFileNames))

		var scanner *bufio.Scanner
		for index, file := range listOfFileNames {

			myFile, err := os.Open(file)
			if err != nil {
				panic(err)
			} // end if

			scanner = bufio.NewScanner(myFile)
			defer myFile.Close()

			for ok := scanner.Scan(); ok; ok = scanner.Scan() {
				dataLines[index] = scanner.Text()
			} // end if

		} // end for loop

		return dataLines

	} // end else

	fmt.Println("The data has been successfully compiled!")
	return dataLines

} // end function

func getCompleteFileNames(baseDir string, partialFileNames []string) []string {

	names := make([]string, len(partialFileNames))

	for index := range partialFileNames {
		names[index] = baseDir + "/" + partialFileNames[index]
	} // end for loop

	return names

} // end func
