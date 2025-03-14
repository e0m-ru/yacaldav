package calDavClient

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/emersion/go-ical"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
	"github.com/google/uuid"
)

var (
	C                = config.LoadConifg() //Config
	dateFormatString = "2006-01-02"
)

func NewCalDavClient() (*caldav.Client, error) {
	client, err := caldav.NewClient(
		webdav.HTTPClientWithBasicAuth(nil,
			C.YaAuth.YAUSER,
			C.YaAuth.CALPWD),
		C.YaAuth.YACAL)
	if err != nil {
		return client, err
	}
	return client, err
}

func GetCalendarsList(client *caldav.Client) (calendars []caldav.Calendar, err error) {
	ctx := context.Background()
	principal, err := client.FindCurrentUserPrincipal(ctx)
	if err != nil {
		return calendars, err
	}
	homeset, err := client.FindCalendarHomeSet(ctx, principal)
	if err != nil {
		return calendars, err
	}
	calendars, err = client.FindCalendars(ctx, homeset)
	if err != nil {
		return calendars, err
	}
	return calendars, err
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

func NewEvent(summ, desc, loc string, st, et time.Time) *ical.Event {
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
	return event
}

func NewCalendar(event *ical.Event) *ical.Calendar {
	cal := ical.NewCalendar()
	cal.Props.SetText(ical.PropVersion, "2.0")
	cal.Props.SetText(ical.PropProductID, C.ProductID)
	cal.Children = append(cal.Children, event.Component)
	var buf bytes.Buffer
	if err := ical.NewEncoder(&buf).Encode(cal); err != nil {
		log.Fatal(err)
	}
	return cal
}
