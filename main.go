package main

import (
	"context"
	"fmt"
	"time"

	calDavClient "github.com/e0m-ru/yacaldav/client"
	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/e0m-ru/yacaldav/report"
	"github.com/emersion/go-webdav/caldav"
)

const (
	DateTime = "2006-01-02"
)

var (
	C                = config.LoadConifg()
	L                = logger.NewLogger(logger.LogLevel(C.Logging.Level), C.Logging.File)
	dateFormatString = "2006-01-02"
)

func main() {
	client, err := calDavClient.NewCalDavClient(C.YaAuth.YAUSER, C.YaAuth.CALPWD, C.YaAuth.YACAL)
	L.Error(err)
	// ctx := context.Background()

	// month, err := time.Parse(dateFormatString, "2025-03-01")
	// L.Error(err)

	// c, err := client.QueryCalendar(ctx, "/calendars/e0m.ru@ya.ru/events-29358211/", calDavClient.BuildMonthRangeQuery(month))
	// L.Error(err)

	// report.PrintAllCalendarsData(c)

	// ev := calDavClient.NewEvent("Название", "Описание", "114", time.Now(), time.Now().Add(time.Hour))
	// cal := calDavClient.NewCalendar(ev)
	// client.PutCalendarObject(ctx, "/calendars/e0m.ru@ya.ru/events-29358211/assa.ics", cal)

	MonthReport(client, 3)
}

func MonthReport(client *caldav.Client, month time.Month) {
	ctx := context.Background()
	now := time.Now()
	year := now.Year()
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lst, err := calDavClient.GetCalendarsList(client)
	L.Error(err)
	for _, calendar := range lst {
		fmt.Fprintf(L.Logger.Writer(), ">>>>>>>>> %v\n  <<<<<<<<<<<<", calendar.Name)
		claList, err := client.QueryCalendar(ctx, calendar.Path, calDavClient.BuildMonthRangeQuery(date))
		L.Error(err)
		report.PrintAllCalendarsData(L.Logger.Writer(), claList)
		fmt.Fprintf(L.Logger.Writer(), " =========================\n\n")
	}
}
