package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type SlackConfig struct {
	ClientID      string `json:"clientID"`      //App Client ID
	ClientSecret  string `json:"clientSecret,omitempty"`  //App Client Secret
	SigningSecret string `json:"signingSecret,omitempty"` //App Signing Secret
	Port          string `json:"port,omitempty"`          //Host Port to connect to
	Host          string `json:"host,omitempty"`          //Host to connect to
}
type PomodoroConfig struct {
	WorkTime     int64  `json:"worktime,omitempty"`  //time to work in pomodoro (25min default)
	Cyles        int    `json:"cycles,omitempty"`    //Number of pomodoro cyles until 
}

type Config struct {
	Slack    SlackConfig
	Pomodoro PomodoroConfig
}

func getSlackConfig(fileContents []byte) SlackConfig {
	var slack SlackConfig = SlackConfig{"client","secret","signing","80","app.slack.com"}
	err := json.Unmarshal(fileContents, &slack)
	if err != nil {
		log.Fatal(err)
	}
	return slack
}

func getPomodoroConfig(fileContents []byte) PomodoroConfig {
	var pomodoro PomodoroConfig = PomodoroConfig{25,4}
	err := json.Unmarshal(fileContents, &pomodoro)
	if err != nil {
		log.Fatal(err)
	}
	return pomodoro
}
func GetAppConfig(configFile string) (Config, error) {

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var config Config = Config{getSlackConfig(fileContents),
		getPomodoroConfig(fileContents)}
	return config, err
}

