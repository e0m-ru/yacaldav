package main

import (
	"context"
	"time"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/e0m-ru/yacaldav/webDavClient"
)

const (
	DateTime = "2006-01-02"
)

var (
	L                = logger.NewLogger(logger.DEBUG, "file")
	C                = config.LoadConifg()
	ctx              = context.Background()
	dateFormatString = "2006-01-02"
)

func main() {

	client := webDavClient.NewCalDavClient()
	month, err := time.Parse(dateFormatString, "2025-03-01")
	c, err := client.QueryCalendar(ctx, "/calendars/e0m.ru@ya.ru/events-29358211/", webDavClient.BuildMonthRangeQuery(month))
	L.Error(err)
	webDavClient.PrintData(c)

	ev := webDavClient.NewEvent("Название", "Описание", "114", time.Now(), time.Now().Add(time.Hour))
	client.PutCalendarObject(ctx, "/calendars/e0m.ru@ya.ru/events-29358211/assa.ics", ev)

	// MonthReport(client)

}
