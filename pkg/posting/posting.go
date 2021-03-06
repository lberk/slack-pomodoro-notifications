package posting

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lberk/sn/pkg/config"
)

func Post(conf config.Config) {
	reqUrl := fmt.Sprintf("%s/api/dnd.setSnooze?token=%s&num_minutes=%d", conf.Slack.Host, html.EscapeString(conf.Slack.Token), conf.Pomodoro.WorkTime)
	response, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body) //Should parse this out a bit better
	reqUrl = fmt.Sprintf("%s/api/users.profile.set", conf.Slack.Host)
	profile, err := json.Marshal(
		&struct {
			StatusText       string `json:"status_text"`
			StatusEmoji      string `json:"status_emoji"`
			StatusExpiration int64  `json:"status_expiration"`
		}{
			StatusText:       conf.Pomodoro.WorkMessage,
			StatusEmoji:      conf.Pomodoro.WorkEmoji,
			StatusExpiration: (time.Now().Unix() + int64(60*conf.Pomodoro.WorkTime)),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	values := url.Values{
		"token":   []string{conf.Slack.Token},
		"profile": []string{string(profile)},
	}
	response, err = http.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	body, err = ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body) //Should parse this out a bit better
}
