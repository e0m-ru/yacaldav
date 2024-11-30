package main

import (
	"context"
	"fmt"
	"os"

	ical "github.com/arran4/golang-ical"
	yaclient "github.com/e0m-ru/yacaldav/client"
	config "github.com/e0m-ru/yacaldav/conf"
	"github.com/emersion/go-webdav/caldav"
	// "github.com/emersion/go-webdav/caldav"
)

func main() {
	ctx := context.Background()
	client := yaclient.NewClient()
	fmt.Printf("%v\n", client)
	s, err := client.Create(ctx, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", s)

	// cal, err := readIcsFromFile("zqp2XHMuyandex.ru.ics")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, c := range cal.Events() {
	// 	fmt.Printf("%v\n", c.Id())
	// 	for i, e := range c.Properties {
	// 		fmt.Printf("% 2v: %-11v %-10v %v\n", i, e.IANAToken, e.GetValueType(), e.Value)
	// 	}
	// }
	// var _ ical.ComponentBase

	// var x ical.VEvent

	// ctx := context.Background()
	// calClient := client.NewClient()

	// printCalendars(ctx, calClient)

	// curl -X PROPFIND -u USERNAME:PASSWORD https://caldav.fastmail.com/dav/principals/user/USERNAME
	// aaa, err := calClient.FindCalendars(ctx, "/calendars/e0m.ru@ya.ru/")

	// aa, err := calClient.Create(ctx, "ASSA")
	// _, err = aa.Write([]byte("ASSA"))
	// aa.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// homeSet, err := calClient.GetCalendarObject(ctx, "/calendars/e0m.ru@yandex.ru/events-12404324/zqp2XHMuyandex.ru.ics")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(homeSet)
}

// ls calendars
func printCalendars(ctx context.Context, calClient *caldav.Client) {
	cfg := config.New()
	fi, err := calClient.FindCalendars(ctx, "/calendars/events-12404324/"+cfg.YaAuth.YAUSER)
	if err != nil {
		panic(err)
	}
	for i, cal := range fi {
		if cal.SupportedComponentSet[0] != "VEVENT" {
			continue
		}
		fmt.Printf("%3d: %-31s %#v\n", i, cal.Name, cal.Path)
		x, err := calClient.GetCalendarObject(ctx, cal.Path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", x)
	}
}

func readIcsFromFile(fileName string) (cal *ical.Calendar, err error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	defer f.Close()
	if err != nil {
		return
	}
	cal, err = ical.ParseCalendar(f)
	if err != nil {
		return
	}
	return
}
