package model

import "time"

type Competition struct {
	Name string

	Location Location

	StartDate time.Time
	EndDate   time.Time

	Rounds []string

	Organiser string
	Phone     string
	Email     string
	Web       string
	Notes     string
}
