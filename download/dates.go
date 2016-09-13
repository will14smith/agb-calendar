package main

import (
	"fmt"
	"time"
)

func monthsBetween(start, end time.Time) []time.Time {
	if !start.Before(end) {
		panic("start is after end")
	}

	var months []time.Time

	currentMonth := start.Month()
	currentYear := start.Year()

	endMonth := end.Month()
	endYear := end.Year()

	for currentYear != endYear || currentMonth != endMonth {
		month, _ := time.Parse("2006-01", fmt.Sprintf("%04d-%02d", currentYear, currentMonth))

		months = append(months, month)

		currentMonth++
		if currentMonth > 12 {
			currentMonth = 1
			currentYear++
		}
	}

	return months
}
