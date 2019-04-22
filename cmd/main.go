package main

import (
	"flag"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"io/ioutil"
)

type slackConfig struct {
	Token     string `json:"token"`//Token to Authorize App
	WorkTime  int64  //Amount of time to work in pomodoro (25 minutes default)
	BreakTime int64  //Amount of time to break after a pomodoro (5 minutes default)
	Cyles     int    //Number of pomodoro cyles until 
}
func main() {

	configFile := flag.String("configFile", "config.json", "Slack Variable config file, default: config.json")
	flag.Parse()
	
	log.Printf("configFile: %s\n", *configFile)

	file, err := os.Open(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var config slackConfig
	err = json.Unmarshal(fileContents, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token from config: %s\n", config.Token)
}
