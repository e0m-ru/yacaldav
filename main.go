package main

import (
	"context"
	"fmt"

	. "github.com/e0m-ru/yacaldav/client"

	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
)

func main() {

	// Инициализация временных диапазонов
	// startTime, _ := time.Parse(time.RFC3339, "2024-01-04T00:00:00Z")
	// endTime, _ := time.Parse(time.RFC3339, "2024-12-05T00:00:00Z")

	// Создаем простейший запрос CalendarQuery.
	query := &caldav.CalendarQuery{
		CompRequest: caldav.CalendarCompRequest{},
		CompFilter: caldav.CompFilter{
			Name: "VEVENT",
		},
	}

	// ---------------CLIENT---------------
	c := webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD)
	client, err := caldav.NewClient(c, C.YaAuth.YACAL)
	L.Error(err)
	ctx := context.Background()

	fmt.Printf("query: %v\n", query)
	// var dd caldav.CalendarMultiGet
	aa, err := client.QueryCalendar(ctx, "/calendars/e0m.ru@ya.ru/", query)
	L.Error(err)
	fmt.Printf("%v\n", aa)

	// principal, err := client.FindCurrentUserPrincipal(ctx)
	// L.Error(err)

	// HomeSet, err := client.FindCalendarHomeSet(ctx, principal)
	// L.Error(err)

	// calendars, err := client.FindCalendars(ctx, HomeSet)
	// L.Error(err)
	// fmt.Printf("%v\n", calendars)

	// var wg sync.WaitGroup
	// for _, calendar := range calendars {
	// 	wg.Add(1)
	// 	go func(wq *sync.WaitGroup, calendar *caldav.Calendar) {
	// 		c := webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD)
	// 		client, err := caldav.NewClient(c, C.YaAuth.YACAL)
	// 		L.Error(err)
	// 		ctx := context.Background()
	// 		stat, err := client.QueryCalendar(ctx, calendar.Path, &caldav.CalendarQuery{})
	// 		L.Error(err)
	// 		fmt.Printf("%v\n", stat)
	// 		wg.Done()
	// 	}(&wg, &calendar)
	// }

	// wg.Wait()
}
