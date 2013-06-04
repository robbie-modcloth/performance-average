package main

import (
	"fmt"
	"os"
)

func function1(data_directory *os.File) {

	fmt.Println((*data_directory).Readdirnames(-1))

} // end function
