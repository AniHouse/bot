package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/anihouse/bot/config"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

type modules map[string]Module

// build variables
var (
	Version string
)

var (
	logger *logrus.Logger
)

// bot instanses
var (
	Session *discordgo.Session

	Cron    = new(cron.Cron)
	Modules = make(modules)
)

func Init() {
	auth()
}

func auth() {
	fmt.Println("Authorization...")

	var err error

	Session, err = discordgo.New(config.Session.Token)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = Logger(config.Session.Log)
	if err != nil {
		log.Fatal(err)
	}

	switch logger.Level {
	case logrus.ErrorLevel:
		Session.LogLevel = discordgo.LogError
	case logrus.WarnLevel:
		Session.LogLevel = discordgo.LogWarning
	case logrus.InfoLevel:
		Session.LogLevel = discordgo.LogInformational
	case logrus.DebugLevel:
		Session.LogLevel = discordgo.LogDebug
	}

	discordgo.Logger = func(msgL, caller int, format string, a ...interface{}) {
		switch msgL {
		case discordgo.LogError:
			logger.Errorf(format, a...)
		case discordgo.LogWarning:
			logger.Warnf(format, a...)
		case discordgo.LogInformational:
			logger.Infof(format, a...)
		case discordgo.LogDebug:
			logger.Debugf(format, a...)
		}
	}

	err = Session.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running as", Session.State.User, "at", time.Now().Format("02-01-2006 15:04:05.999999999"))
}
