package main

import (
	"generateEventData/config"
	"generateEventData/services"
	"log"
)

func main() {
	log.Printf("starting generating data in %s format", config.Config.OutputFormat)
	log.Printf("Generating %d number of events", config.Config.DataSetSize)
	if config.Config.OutputFormat == "csv" {
		services.GenerateCSVDataV2MemConserve()
		//services.GenerateCSVDataV1()
	} else if config.Config.OutputFormat == "json" {
		services.GenerateJSONData()
	} else {
		panic("cant generate something else")
	}

}
