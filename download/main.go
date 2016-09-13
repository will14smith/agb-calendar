package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

var dataRoot = path.Join("..", "data")

const monthCalendarUrl = "http://www.archerygb.org/tools/cache/calendar/ENG-MkI3RUQxQzA0NTQxNEVCRkExNTg4NzQ0Q0YyNEZBNjI=-83a6e51afa953dce4c2adae448662fef?period=%s"
const dayCalendarUrl = monthCalendarUrl + "&seldate=%s"

func main() {
	// hard coded dates won't have errors
	startDate, _ := time.Parse("2006-01-02", "2016-10-01")
	endDate, _ := time.Parse("2006-01-02", "2017-10-01")

	// get the months to look through
	months := monthsBetween(startDate, endDate)

	for _, month := range months {
		fmt.Println("Processing", month.Format("2006-01"))

		// get month calendar
		node, err := download(fmt.Sprintf(monthCalendarUrl, month.Format("2006-01")))
		if err != nil {
			panic(err)
		}

		var monthRoot = path.Join(dataRoot, month.Format("2006-01"))
		err = os.Mkdir(monthRoot, os.ModeDir)
		if err != nil {
			panic(err)
		}

		// save page
		file, err := os.Create(path.Join(monthRoot, "index.html"))
		if err != nil {
			panic(err)
		}
		html.Render(file, node)
		file.Close()

		// get days
		days := getDaysFromMonth(node)
		fmt.Println(" ==> Starting", len(days), "days")

		for _, day := range days {
			node, err := download(fmt.Sprintf(dayCalendarUrl, month.Format("2006-01"), day.Format("2006-01-02")))

			// save page
			file, err := os.Create(path.Join(monthRoot, fmt.Sprintf("%02d.html", day.Day())))
			if err != nil {
				panic(err)
			}
			html.Render(file, node)
			file.Close()
		}
	}
}

func getDaysFromMonth(node *html.Node) []time.Time {
	var days []time.Time

	eventDays := scrape.FindAll(node, scrape.ByClass("eventday"))
	for _, eventDay := range eventDays {
		id := scrape.Attr(eventDay, "id")

		day, err := time.Parse("2006-01-02", id[len("day-"):len("day-2006-01-02")])
		if err != nil {
			panic(err)
		}

		days = append(days, day)
	}

	return days
}
