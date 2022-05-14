package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(LearnWeekdayofBirth(1983, time.July, 10))
	// output : Sunday :))
}

func LearnWeekdayofBirth(year int, month time.Month, day int) time.Weekday {
	getDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return getDay.Weekday()
}
