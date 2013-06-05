package main

import (

  "strings"

)

type Datum struct {
  date string
  execution_time_in_seconds float32
  execution_time_and_latency float32
  errors string
}


func parse_data(data_string string) {
  
  string_array := strings.Split(data_string, "\n")

  data := make([]Datum, len(string_array))

  // now string_array is an array of each line, and therefore each piece of datum
  for line := range string_array {
    
  }
  
} // end function

