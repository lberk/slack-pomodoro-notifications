package posting

import (
	"fmt"
	"github.com/slack-pomodoro-notifications/internal/config"
)

func PostingTest(conf config.Config) {
	fmt.Println("HERE!", conf.Slack.ClientID)
}

