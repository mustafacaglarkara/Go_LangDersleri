package main

import (
	"fmt"
	"time"
)

func main() {
	//call the function
	fmt.Println(LearnWeekdayofBirth(1983, time.July, 10))
	// output : Sunday :))
}

// LearnWeekdayofBirth create function to find which day of the week we were born on
func LearnWeekdayofBirth(year int, month time.Month, day int) time.Weekday {
	// define variable and give the variable as a parameter
	getDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	// return the day of the week that we born on
	return getDay.Weekday()
}
