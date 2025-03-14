package report

import (
	"fmt"
	"io"
	"time"

	"github.com/emersion/go-ical"
	"github.com/emersion/go-webdav/caldav"
)

func printEvent(cal ical.Event) string {
	startTime, _ := cal.DateTimeStart(time.Local)
	endTime, _ := cal.DateTimeEnd(time.Local)
	html := fmt.Sprintf("<p>start date : %s</p>\n<p>start time: %s</p>\n",
		startTime.Format("2006-01-02"), startTime.Format("15:04"))
	sy, sm, sd := startTime.Date()
	ey, em, ed := endTime.Date()
	if sy != ey || sm != em || sd != ed {
		html += fmt.Sprintf("<p>end date: %s</p>", endTime.Format("2006-01-02"))
	}
	html += fmt.Sprintf("<p>end time: %s</p>\n", endTime.Format("15:04"))
	title := getPropText(cal, ical.PropSummary)
	desc := getPropText(cal, ical.PropDescription)
	loc := getPropText(cal, ical.PropLocation)
	// dep := getPropText(cal, "DEPARTMENT")
	uid := getPropText(cal, ical.PropUID)
	return html + fmt.Sprintf("uid: %s\n<h1>%s</h1>\n<p>Location: %s</p>\n<p>%s</p>\n", uid, title, loc, desc)
}

func getPropText(cal ical.Event, propName string) string {
	prop := cal.Props.Get(propName)
	if prop != nil {
		text, _ := prop.Text()
		return text
	}
	return ""
}

func PrintAllCalendarsData(w io.Writer, calendarList []caldav.CalendarObject) {
	for _, c := range calendarList {
		fmt.Fprintf(w, "path: %s\n", c.Path)
		fmt.Fprintf(w, "modTime %s\n", c.ModTime)
		fmt.Fprintf(w, "---------Events--------\n")
		for _, e := range c.Data.Events() {
			fmt.Fprintf(w, "%s\n-------------------------\n", printEvent(e))
		}
	}
}
