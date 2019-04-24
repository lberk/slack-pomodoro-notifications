package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/slack-pomodoro-notifications/internal/config"
	"github.com/slack-pomodoro-notifications/internal/posting"
)

func main() {

	configFile := flag.String("configFile", "config.json", "Slack Variable config file, default: config.json")
	flag.Parse()
	log.Printf("configFile: %s\n", *configFile)

	config, err := config.GetAppConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("token from config: ", config)
	posting.PostingTest(config)
}
