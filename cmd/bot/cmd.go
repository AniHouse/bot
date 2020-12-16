package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/anihouse/bot"
	"github.com/anihouse/bot/app"
	"github.com/anihouse/bot/config"
	"github.com/anihouse/bot/db"
	"github.com/anihouse/bot/tpl"
	"github.com/urfave/cli"
)

func run(c *cli.Context) error {
	fmt.Println("Bot is running. Press Ctrl + C to exit.")

	config.Load(c.GlobalString("config"))

	tpl.Init()
	db.Init()
	bot.Init()
	app.Init()

	bot.Modules.Init()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}
