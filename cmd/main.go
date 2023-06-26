package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"main.go/payload"
	"main.go/user"
)

var Filename = flag.String("config", "../conf/config.json", "Location of the config file.")

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	config, err := loadConfig(*Filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Available configuration :  %+v \n", config)

	defer wg.Wait()
	for {
		wg.Add(1)
		go user.GetUserDetails(wg, mu, config)
		time.Sleep(time.Minute * time.Duration(config.DataFetchDurationInMinute))
	}
}

// loadConfig function load configuration of project
func loadConfig(fileName string) (payload.Config, error) {
	config := payload.Config{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return config, fmt.Errorf("Error while reading config file %v", err.Error())
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Error while unmarshling config file %v", err.Error())
	}

	return config, nil
}
