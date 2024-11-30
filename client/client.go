package client

import (
	config "github.com/e0m-ru/yacaldav/conf"
	"github.com/e0m-ru/yacaldav/logger"
	"github.com/emersion/go-webdav"

	// "github.com/emersion/go-webdav/caldav"
	"github.com/trvita/caldav-client-yandex/caldav"
)

var L logger.Logger

func init() {
	L, err := logger.NewLogger(logger.ERROR, "")
	if err != nil {
		L.Error("can't launch Loger")
	}
	L.Info("Logger started")
}

func NewClient() (client *caldav.Client) {
	cfg := config.New()
	c := webdav.HTTPClientWithBasicAuth(nil, cfg.YaAuth.YAUSER, cfg.YaAuth.CALPWD)
	client, err := caldav.NewClient(c, cfg.YaAuth.YACAL) // cfg.YaAuth.YACAL+"/"+cfg.YaAuth.YACAL)
	if err != nil {
		panic(err)
	}
	return
}
