package main

import (
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"mofu/core"
	"mofu/ent"
	"mofu/tg"
	"mofu/tw"
	"mofu/web"
	"os"
	"strconv"
)

var (
	TelegramBotToken   = os.Getenv("TELEGRAM_BOT_TOKEN")
	ControlRoomID, _   = strconv.ParseInt(os.Getenv("CONTROL_ROOM_ID"), 10, 64)
	BroadcastRoomID, _ = strconv.ParseInt(os.Getenv("BROADCAST_ROOM_ID"), 10, 64)
	WarningRoomID, _   = strconv.ParseInt(os.Getenv("WARNING_ROOM_ID"), 10, 64)
	TwitterApiToken    = os.Getenv("TWITTER_API_TOKEN")
	Dev                = os.Getenv("Dev") == "1"
	PORT               = os.Getenv("PORT")
	Database           = os.Getenv("DATABASE")
)

func init() {
}

func main() {
	if !Dev {
		log.SetLevel(log.WarnLevel)

	}
	create()
	<-make(chan int)
}

func create() {
	client, err := ent.Open(dialect.Postgres, Database)
	if err != nil {
		panic(err)
	}
	//client = client.Debug()
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	twitter := tw.New(TwitterApiToken)
	c := core.New(client, twitter)
	telegram := tg.New(TelegramBotToken, ControlRoomID, BroadcastRoomID, WarningRoomID, c)
	if !Dev {
		log.SetOutput(telegram)
	}

	w := web.New(c)

	go telegram.RunTask()
	go telegram.Update()
	go telegram.UpdateLog()
	go c.UpdateSubscribe()
	go c.Update()
	if PORT != "" {
		go w.Run(PORT)
	} else {
		go w.Run()
	}

}
