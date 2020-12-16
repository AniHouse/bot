package main

import (
	"log"
	"os"

	_ "github.com/anihouse/bot/modules/clubs"
	_ "github.com/anihouse/bot/modules/test"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	if err := application().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
