package main

import (
	"flag"
	"github.com/slack-pomodoro-notifications/internal/config"
	"github.com/slack-pomodoro-notifications/internal/posting"
	"log"
)

func main() {

	configFile := flag.String("configFile", "config.json", "Slack Variable config file, default: config.json")
	flag.Parse()
	log.Printf("configFile: %s\n", *configFile)

	config, err := config.GetAppConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	posting.Post(config)
}
