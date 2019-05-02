package posting

import (
	"fmt"
	"github.com/lberk/sn/internal/config"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func Post(conf config.Config) {
	url := fmt.Sprintf("%s/api/dnd.setSnooze?token=%s&num_minutes=%d", conf.Slack.Host, html.EscapeString(conf.Slack.Token), conf.Pomodoro.WorkTime)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body) //Should parse this out a bit better
}
