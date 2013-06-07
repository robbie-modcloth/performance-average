package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
  "bytes"
)

func compileTheData(dataDirectory *os.File)  {


  var total float64 = 0.0
  var totalWLatency float64 = 0.0
  var length int = 0
  var numExecutionsCompleted = 0

  var tempTotal, tempTotalWLatency float64
  var tempLength, tempNumExecutionsCompleted int

	if partialFileNames, err := dataDirectory.Readdirnames(-1); err != nil {
		panic(err)
	} else {

		baseDir := dataDirectory.Name()
		listOfFileNames := getCompleteFileNames(baseDir, partialFileNames)
    fileChannel, numFiles := getAllFiles(listOfFileNames)

		for file := range fileChannel {

      f = func(file *os.File) (total, totalWLatency float64, length int, completion int ) {
        total, totalWLatency, length = calcPerfNumbers(file)
        return total, totalWLatency, length, 1
      }

      tempTotal, tempTotalWLatency, tempLength, tempNumExecutionsCompleted = go f(file)
      
      total, totalWLatency += tempTotal, tempTotalWLatency
      

		} // end for loop

	} // end else

	fmt.Println("The data has been successfully compiled!") 

} // end function

func getCompleteFileNames(baseDir string, partialFileNames []string) []string {

	names := make([]string, len(partialFileNames))

	for index := range partialFileNames {
		names[index] = baseDir + "/" + partialFileNames[index]
	} // end for loop

	return names

} // end func

func getAllFiles(name_slice []string) (chan *os.File, int) {

  fileChannel := make(chan *os.File, 1000)
  numFiles := 0

  for _, name := range name_slice {

    if file, err := os.Open(name); err != nil {
      panic(err)
    } else {
      fileChannel <- file
    }
    numFiles++

  } // end for loop

  return fileChannel, numFiles
  
} // end func

func calcPerfNumbers(file *os.File) (total float64, totalWLatency float64, numLines int) {
  
  scanner := bufio.NewScanner(file)
  txt := ""
  numLines = 0
  total, totalWLatency = 0.0, 0.0
  defer file.Close()


  var val1, val2 float64
  for ok := scanner.Scan() ; ok ; ok = scanner.Scan() {
    
    txt = scanner.Text()
    val1, val2 = parseLine(txt)
    total += val1
    totalWLatency += val2
    
    numLines++
    
  } // end for loop
  
  
  return total, totalWLatency, numLines
  
} // end func

func parseLine(str string) (time, timeWLatency float64) {
  
  reader := strings.NewReader(str)
  scanner := bufio.NewScanner(reader)
  scanner.Split(ScanLinesForComma)
  
  dataSlice := make([]string, 4)
  
  for ok, index := scanner.Scan(), 0  ; ok ; ok = scanner.Scan() {
    dataSlice[index] = scanner.Text()
    index++
  } // end for loop

  result1, _ := strconv.ParseFloat(dataSlice[1], 64)
  result2, _ := strconv.ParseFloat(dataSlice[2], 64)

  return result1, result2
} // end func


func ScanLinesForComma(data []byte, atEOF bool)    (advance int, token []byte, err error) {
  if atEOF && len(data)     == 0 {
    return 0, nil, nil
}
  if i := bytes.IndexByte(data, ','); i >= 0 {
  // We have a full newline-terminated line.
    return i + 1, data[0:i], nil
  }
  // If we're at EOF, we have a final, non-terminated line. Return it.
  if atEOF {
    return len(data), data, nil
  }
  // Request more data.
  return 0, nil, nil
}

