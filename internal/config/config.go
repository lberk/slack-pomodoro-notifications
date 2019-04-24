package config

import (
	"log"
	"os"
)

type SlackConfig struct {
	ClientID      string `json:"clientID,omitempty"`      //App Client ID
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

func GetAppConfig(configFile string) (Config, error) {

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

//	fileContents, err := ioutil.ReadAll(file)
//	if err != nil {
//		log.Fatal(err)
//	}
	pomodoro := PomodoroConfig{25, 4}
	slack    := SlackConfig{"foobar","hello","world","80","app.slack.com"}
	config   := Config{slack,pomodoro}
	return config, nil
	
}

//func GetSlackConfig(r io.Reader) (Config, error) {
//}
