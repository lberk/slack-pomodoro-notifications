package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type SlackConfig struct {
	Token string `json:"token,omitempty"` //OAuth Token
	Host  string `json:"host,omitempty"`  //Host to connect to
}
type PomodoroConfig struct {
	WorkTime int64 `json:"worktime,omitempty"` //time to work in pomodoro (25min default)
}

type Config struct {
	Slack    SlackConfig
	Pomodoro PomodoroConfig
}

func getSlackConfig(fileContents []byte) SlackConfig {
	var slack SlackConfig = SlackConfig{"token", "https://slack.com"}
	err := json.Unmarshal(fileContents, &slack)
	if err != nil {
		log.Fatal(err)
	}
	return slack
}

func getPomodoroConfig(fileContents []byte) PomodoroConfig {
	var pomodoro PomodoroConfig = PomodoroConfig{25}
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
