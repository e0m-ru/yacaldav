package client

import (
	"context"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
)

var (
	L   *logger.Logger
	C   *config.Config
	ctx = context.Background()
)

func init() {
	L = logger.NewLogger(logger.DEBUG, "")
	C = config.LoadConifg()
}

func NewClient() (client *caldav.Client, err error) {
	c := webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD)
	client, err = caldav.NewClient(c, C.YaAuth.YACAL)
	L.Error(err)
	return
}

func CalendarsSet(client *caldav.Client) *[]caldav.Calendar {
	principal, err := client.FindCurrentUserPrincipal(ctx)
	L.Error(err)

	HomeSet, err := client.FindCalendarHomeSet(ctx, principal)
	L.Error(err)

	set, err := client.FindCalendars(ctx, HomeSet)
	L.Error(err)

	return &set
}
