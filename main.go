package yacaldav

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-ical"
	"github.com/emersion/go-webdav/caldav"
)

const (
	DateTime = "2006-01-02"
)

var (
	C = config.LoadConifg()
	L = logger.NewLogger(logger.LogLevel(C.Logging.Level), C.Logging.File)
)

func main_temp() {
	client, err := NewCalDavClient(C.YaAuth.YAUSER, C.YaAuth.CALPWD, C.YaAuth.YACAL)
	L.Error(err)
	// ctx := context.Background()

	// month, err := time.Parse(dateFormatString, "2025-03-01")
	// L.Error(err)

	// c, err := client.QueryCalendar(ctx, "/calendars/e0m.ru@ya.ru/events-29358211/", yacaldav.BuildMonthRangeQuery(month))
	// L.Error(err)

	// report.PrintAllCalendarsData(c)

	// ev := yacaldav.NewEvent("Название", "Описание", "114", time.Now(), time.Now().Add(time.Hour))
	// cal := yacaldav.NewCalendar(ev)
	// client.PutCalendarObject(ctx, "/calendars/e0m.ru@ya.ru/events-29358211/assa.ics", cal)

	MonthReport(client, 3)
}

func MonthReport(client *caldav.Client, month time.Month) {
	ctx := context.Background()
	now := time.Now()
	year := now.Year()
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lst, err := GetCalendarsList(client)
	L.Error(err)
	for _, calendar := range lst {
		fmt.Fprintf(L.Logger.Writer(), ">>>>>>>>> %v\n  <<<<<<<<<<<<", calendar.Name)
		claList, err := client.QueryCalendar(ctx, calendar.Path, BuildMonthRangeQuery(date))
		L.Error(err)
		PrintAllCalendarsData(L.Logger.Writer(), claList)
		fmt.Fprintf(L.Logger.Writer(), " =========================\n\n")
	}
}

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
