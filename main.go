package main

import (
	"context"
	"fmt"

	"github.com/e0m-ru/yacaldav/client"
	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
)

var L *logger.Logger

func init() {
	L, _ = logger.NewLogger(logger.DEBUG, "")
}
func main() {
	C := config.New()
	// crete empty content
	var ctx = context.Background()
	// add timeout to context default 1000
	ctx, _ = context.WithTimeout(ctx, C.Net.Timeout)

	// create new CalDAV client
	cc, err := client.NewClient()
	L.Error(err)

	// find Prinxipal string
	// /principals/users/e0m.ru@ya.ru/
	P, err := cc.FindCurrentUserPrincipal(ctx)
	L.Error(err)
	fmt.Printf("Principal: %q\n", P)

	// find HomeSet
	// /calendars/e0m.ru@ya.ru/
	HS, err := cc.FindCalendarHomeSet(ctx, P)
	L.Error(err)
	fmt.Printf("HomeSet: %q\n", HS)

	// get calendars
	calSet, err := cc.FindCalendars(ctx, HS)
	L.Error(err)
	// print calendars
	for i, c := range calSet {
		fmt.Printf("%02v:%-6v %v %q\n", i, c.SupportedComponentSet, c.Path, c.Name)
	}

	// file info
	fi, err := cc.Stat(ctx, "calendars/e0m.ru@ya.ru/events-12404324/")
	L.Error(err)
	fmt.Printf("%#v\n", fi)

	// a, err := client.GetCalendarObject(ctx, "calendars/e0m.ru@ya.ru/events-12404324/7jjadJP3yandex.ru.ics")
	// L.Error(err)

	// fmt.Printf("%v\n", a.ContentLength)
	// x := a.Data.Events()
	// fmt.Printf("%v\n", x[0].Component.Props.Get("SUMMARY").Value)
	// fmt.Printf("%s\n", x[0].Component.Props.Get("DESCRIPTION").Value)
	// for i, p := range x[0].Component.Props {
	// 	fmt.Printf("%15s: %v\n", i, p)
	// }
}
