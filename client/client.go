package client

import (
	"context"

	"github.com/e0m-ru/yacaldav/config"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
)

var L *logger.Logger
var C *config.Config

func init() {
	L, _ = logger.NewLogger(logger.DEBUG, "")
	C = config.New()
}

func NewClient() (client *caldav.Client, err error) {
	c := webdav.HTTPClientWithBasicAuth(nil, C.YaAuth.YAUSER, C.YaAuth.CALPWD)
	client, err = caldav.NewClient(c, C.YaAuth.YACAL)
	L.Error(err)
	return
}

func CalendarsSet(client *caldav.Client) *[]caldav.Calendar {
	ctx := context.Background()
	principal, err := client.FindCurrentUserPrincipal(ctx)
	L.Error(err)

	HomeSet, err := client.FindCalendarHomeSet(ctx, principal)
	L.Error(err)

	set, err := client.FindCalendars(ctx, HomeSet)
	L.Error(err)

	return &set
}
