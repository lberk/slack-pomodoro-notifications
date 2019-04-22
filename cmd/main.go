package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

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
	fmt.Printf("File contents:\n%s", fileContents)
}
