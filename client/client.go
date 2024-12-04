package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	ical "github.com/arran4/golang-ical"
	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-webdav"

	// "github.com/emersion/go-webdav/caldav"
	"github.com/trvita/caldav-client-yandex/caldav"
)

var L logger.Logger

func init() {
	L, err := logger.NewLogger(logger.ERROR, "")
	if err != nil {
		L.Error(errors.New("can't launch Loger"))
	}
	L.Info("Logger started")
}

func NewClient() (client *caldav.Client, err error) {
	cfg := config.New()
	c := webdav.HTTPClientWithBasicAuth(nil, cfg.YaAuth.YAUSER, cfg.YaAuth.CALPWD)
	client, err = caldav.NewClient(c, cfg.YaAuth.YACAL) // cfg.YaAuth.YACAL+"/"+cfg.YaAuth.YACAL)
	if err != nil {
		return nil, err
	}
	return
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

func CalendarsSet() *[]caldav.Calendar {
	L, err := logger.NewLogger(logger.DEBUG, "") // logger
	C := config.New()                            // config
	ctx := context.Background()                  // context
	err = errors.New("all fine")

	// New CalDAV client
	CalDAVClient, err := caldav.NewClient(webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD), C.YaAuth.YACAL)
	L.Error(err)

	principal, err := CalDAVClient.FindCurrentUserPrincipal(ctx)
	L.Error(err)

	HomeSet, err := CalDAVClient.FindCalendarHomeSet(ctx, principal)
	L.Error(err)

	set, err := CalDAVClient.FindCalendars(ctx, HomeSet)
	L.Error(err)

	return &set
}
