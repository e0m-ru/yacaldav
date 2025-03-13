package webDavClient

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-ical"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
	"github.com/google/uuid"
)

const (
	ProductID = "-//ittsc Comp//OTT Calendar Version 1.0//RU"
)

var (
	L                = logger.NewLogger(logger.DEBUG, "file")
	C                = config.LoadConifg()
	ctx              = context.Background()
	dateFormatString = "2006-01-02"
)

func NewWebdavClient() (client *caldav.Client, err error) {
	c := webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD)
	client, err = caldav.NewClient(c, C.YaAuth.YACAL)
	L.Error(err)
	return
}

func PrintData(calendarList []caldav.CalendarObject) {
	for _, c := range calendarList {
		for _, e := range c.Data.Events() {
			fmt.Printf("%s\n\n", WriteHTML(e))
		}
	}
}

func GetCalendarsList(client *caldav.Client) []caldav.Calendar {
	principal, err := client.FindCurrentUserPrincipal(ctx)
	L.Error(err)
	homeset, err := client.FindCalendarHomeSet(ctx, principal)
	L.Error(err)
	calendars, err := client.FindCalendars(ctx, homeset)
	L.Error(err)
	return calendars
}

func BuildMonthRangeQuery(month time.Time) *caldav.CalendarQuery {
	compFilter := caldav.CompFilter{
		Name: "VCALENDAR",
		Props: []caldav.PropFilter{
			{Name: "Name"},
		},
		Comps: []caldav.CompFilter{{
			Name:  "VEVENT",
			Start: month,
			End:   month.AddDate(0, 1, 0),
			Props: []caldav.PropFilter{{
				Name: "SUMMARY",
				TextMatch: &caldav.TextMatch{
					Text: "ОТТ",
				},
			}},
		}},
	}
	query := caldav.CalendarQuery{
		CompFilter: compFilter,
	}
	return &query
}

func WriteHTML(cal ical.Event) string {
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

	return html + fmt.Sprintf("<h1>%s</h1>\n<p>Location: %s</p>\n<p>%s</p>\nuid: %s\n", title, loc, desc, uid)
}

func getPropText(cal ical.Event, propName string) string {
	prop := cal.Props.Get(propName)
	if prop != nil {
		text, _ := prop.Text()
		return text
	}
	return ""
}

func NewCalDavClient() *caldav.Client {
	client, err := caldav.NewClient(
		webdav.HTTPClientWithBasicAuth(nil,
			C.YaAuth.YAUSER,
			C.YaAuth.CALPWD),
		C.YaAuth.YACAL)
	L.Error(err)
	return client
}

func MonthReport(client *caldav.Client) {
	month, err := time.Parse(dateFormatString, "2025-04-01")
	L.Error(err)
	for _, calendar := range GetCalendarsList(client) {
		claList, err := client.QueryCalendar(ctx, calendar.Path, BuildMonthRangeQuery(month))
		L.Error(err)
		PrintData(claList)
	}
}

func NewEvent(summ, desc, loc string, st, et time.Time) *ical.Calendar {
	event := ical.NewEvent()
	uid := uuid.New().String()
	event.Props.SetText(ical.PropUID, uid)
	event.Props.SetDateTime(ical.PropDateTimeStamp, time.Now())
	event.Props.SetText(ical.PropSummary, summ)
	event.Props.SetText(ical.PropDescription, desc)
	event.Props.SetText(ical.PropLocation, loc)
	// event.Props.SetText("DEPARTMENT", "ОТТ")
	event.Props.SetDateTime(ical.PropDateTimeStart, st)
	event.Props.SetDateTime(ical.PropDateTimeEnd, et)

	cal := ical.NewCalendar()
	cal.Props.SetText(ical.PropVersion, "2.0")
	cal.Props.SetText(ical.PropProductID, ProductID)
	cal.Children = append(cal.Children, event.Component)

	var buf bytes.Buffer
	if err := ical.NewEncoder(&buf).Encode(cal); err != nil {
		log.Fatal(err)
	}

	// L.Info(buf.String())
	return cal
}
