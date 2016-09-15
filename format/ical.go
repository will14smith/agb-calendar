package format

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/will14smith/agb-calendar/model"
)

const ICalDateFormat = "20060102"

func ConvertToICal(competitions []*model.Competition) string {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "BEGIN:VCALENDAR")

	for _, competition := range competitions {
		writeCompetitionToICal(&buf, competition)
	}

	fmt.Fprintln(&buf, "END:VCALENDAR")

	return buf.String()
}

func writeCompetitionToICal(writer io.Writer, competition *model.Competition) {
	fmt.Fprintln(writer, "BEGIN:VEVENT")

	fmt.Fprintf(writer, "DTSTART;VALUE=DATE:%s\n", competition.StartDate.Format(ICalDateFormat))
	fmt.Fprintf(writer, "DTEND;VALUE=DATE:%s\n", competition.EndDate.AddDate(0, 0, 1).Format(ICalDateFormat))
	fmt.Fprintf(writer, "SUMMARY:%s\n", competition.Name)
	fmt.Fprintf(writer, "DESCRIPTION:%s\n", getICalDescription(competition))

	fmt.Fprintln(writer, "END:VEVENT")
}

func getICalDescription(competition *model.Competition) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Location: %s\\n", competition.Location.Name)
	fmt.Fprintf(&buf, "Rounds: %s\\n", strings.Join(competition.Rounds, "; "))
	fmt.Fprintf(&buf, "Organiser: %s\\n", competition.Organiser)
	fmt.Fprintf(&buf, "Phone: %s\\n", competition.Phone)
	fmt.Fprintf(&buf, "Email: %s\\n", competition.Email)
	fmt.Fprintf(&buf, "Web: %s\\n", competition.Web)
	fmt.Fprintf(&buf, "Notes: %s\\n", competition.Notes)

	str := buf.String()

	return strings.Replace(str, "\n", "\\n", -1)
}
