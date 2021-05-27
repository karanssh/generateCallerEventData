package config

import (
	"encoding/json"
	"log"
	"os"
)

type (
	configType struct {
		OutputFormat   string
		OutputFileName string
		DataSetSize    int
		DataSetBreaker int
		Type1          int //relative percentage
		Type2          int
		Type3          int
		Type4          int
		RandomTextSize int
		PhoneNoLow     int
		PhoneNoHigh    int
		DurationLow    int
		DurationHigh   int
	}
)

var (

	// Config contains the config items
	Config = configType{
		OutputFormat:   "json",
		OutputFileName: "output",
		DataSetSize:    2000000,
		DataSetBreaker: 1000000,
		Type1:          25,
		Type2:          25,
		Type3:          25,
		Type4:          25,
		RandomTextSize: 7,
		PhoneNoLow:     923152265,
		PhoneNoHigh:    980809969,
		DurationLow:    5,
		DurationHigh:   900,
	}
)

func loadJSON(jsonFile string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// log.Println("path is " + path)
	log.Println("path of config file is " + path + jsonFile)
	f, err := os.Open(path + jsonFile)
	if nil != err {
		log.Printf("Error loading app configuration: err=%v", err)
		return
	}

	dec := json.NewDecoder(f)

	dec.Decode(&Config)
}

func init() {
	loadJSON("/confVolume/config.json")
}
