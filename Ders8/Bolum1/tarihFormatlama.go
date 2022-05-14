package main

import (
	"fmt"
	"time"
)

func main() {
	// get the datetime of now
	simdi := time.Now()
	// define the time format day must be 01 month must be 2 and year 2006
	kisaTarihFormat := "02-01-2006"
	// give the date format style as a parameter to format function
	fmt.Println(simdi.Format(kisaTarihFormat))

	// output is :  05-13-2022
}
