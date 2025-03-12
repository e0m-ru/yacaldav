package main

import (
	"context"
	"fmt"
	"time"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
)

var (
	L        = logger.NewLogger(logger.DEBUG, "")
	C        = config.LoadConifg()
	ctx      = context.Background()
	tl       = "2006-01-02"
	month, _ = time.Parse(tl, "2025-02-01")
)

func main() {

	client, err := caldav.NewClient(
		webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD),
		C.YaAuth.YACAL)
	L.Error(err)

	compFilter := caldav.CompFilter{
		Name: "VCALENDAR",
		Comps: []caldav.CompFilter{{
			Name:  "VEVENT",
			Start: month,
			End:   month.AddDate(0, 1, 0)}},
	}
	query := caldav.CalendarQuery{
		CompFilter: compFilter,
	}
	calendars := getCalendars(client, C.YaAuth.YAUSER, C.YaAuth.CALPWD)

	for _, c := range calendars {
		fmt.Printf("-------------%s-------------\n", c.Name)
		extractDtat(client, query, c.Path)
	}

}

func extractDtat(client *caldav.Client, query caldav.CalendarQuery, calendarUrl string) {
	wdfi, err := client.QueryCalendar(ctx, calendarUrl, &query)
	L.Error(err)
	for _, c := range wdfi {
		fmt.Printf("%s %s\n", c.Data.Events()[0].Component.Props.Get("DTSTART").Value, c.Data.Events()[0].Component.Props.Get("SUMMARY").Value)
	}
}

func getCalendars(client *caldav.Client, yaUser, yaPwd string) []caldav.Calendar {
	principal, err := client.FindCurrentUserPrincipal(ctx)
	L.Error(err)
	fmt.Printf("principal: %s\n", principal)

	homeset, err := client.FindCalendarHomeSet(ctx, principal)
	L.Error(err)
	fmt.Printf("  homeset: %s\n", homeset)

	calendars, err := client.FindCalendars(ctx, homeset)
	L.Error(err)
	fmt.Printf("calendars: \n")
	return calendars
}
