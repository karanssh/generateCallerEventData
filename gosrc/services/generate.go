package services

import (
	"encoding/json"
	"fmt"
	"generateEventData/config"
	"generateEventData/datadefinition"
	"generateEventData/eventgen"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

var (
	totalItems     = config.Config.DataSetSize
	typeOneItems   int
	typeTwoItems   int
	typeThreeItems int
	typeFourItems  int
)

func init() {
	typeOneItems = (config.Config.Type1 * totalItems) / 100
	typeTwoItems = (config.Config.Type2 * totalItems) / 100
	typeThreeItems = (config.Config.Type3 * totalItems) / 100
	typeFourItems = (config.Config.Type4 * totalItems) / 100
	log.Printf("\n INIT \n  Type 1 items %d \n Type 2 items %d \n Type 3 items %d \n Type 4 items %d \n ", typeOneItems, typeTwoItems, typeThreeItems, typeFourItems)
}

func generateItemData() []datadefinition.Event {
	log.Printf("\n Type 1 items %d \n Type 2 items %d \n Type 3 items %d \n Type 4 items %d \n ", typeOneItems, typeTwoItems, typeThreeItems, typeFourItems)
	size := typeOneItems + typeTwoItems + typeThreeItems + typeFourItems
	events := make([]datadefinition.Event, size)
	for i := 0; i < typeOneItems; i++ {
		events[i] = eventgen.GenerateEvent(1)
	}
	for i := typeOneItems; i < typeOneItems+typeTwoItems; i++ {
		events[i] = eventgen.GenerateEvent(2)
	}
	for i := typeOneItems + typeTwoItems; i < typeOneItems+typeTwoItems+typeThreeItems; i++ {
		events[i] = eventgen.GenerateEvent(3)
	}
	for i := typeOneItems + typeTwoItems + typeThreeItems; i < size; i++ {
		events[i] = eventgen.GenerateEvent(4)
	}

	return events
}

func generateItemDataWithCount(typeOneItems, typeTwoItems, typeThreeItems, typeFourItems int) []datadefinition.Event {
	log.Printf("\n Type 1 items %d \n Type 2 items %d \n Type 3 items %d \n Type 4 items %d \n ", typeOneItems, typeTwoItems, typeThreeItems, typeFourItems)
	events := eventgen.GenerateEvents(typeOneItems, 1)
	events = append(events, eventgen.GenerateEvents(typeTwoItems, 2)...)
	events = append(events, eventgen.GenerateEvents(typeThreeItems, 3)...)
	events = append(events, eventgen.GenerateEvents(typeFourItems, 4)...)
	return events
}

func generateItemDataWithCountNewImpl(typeOneItems, typeTwoItems, typeThreeItems, typeFourItems int) []datadefinition.Event {
	generateItemDataWithCountNewImplTimeStart := time.Now()
	log.Printf("generateItemDataWithCountNewImpl started at %s", time.Now().String())
	size := typeOneItems + typeTwoItems + typeThreeItems + typeFourItems
	events := make([]datadefinition.Event, size)
	for i := 0; i < typeOneItems; i++ {
		events[i] = eventgen.GenerateEvent(1)
	}
	for i := typeOneItems; i < typeOneItems+typeTwoItems; i++ {
		events[i] = eventgen.GenerateEvent(2)
	}
	for i := typeOneItems + typeTwoItems; i < typeOneItems+typeTwoItems+typeThreeItems; i++ {
		events[i] = eventgen.GenerateEvent(3)
	}
	for i := typeOneItems + typeTwoItems + typeThreeItems; i < size; i++ {
		events[i] = eventgen.GenerateEvent(4)
	}
	log.Printf("generateItemDataWithCountNewImpl took %s time", time.Since(generateItemDataWithCountNewImplTimeStart).String())
	return events
}

func GenerateJSONData() error {

	data := generateItemData()
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatalf("json.Marshal() err: %v", err)
	}
	GenerateFile(b)
	return nil
}

func GenerateCSVDataV1() error {
	generateCSVDataV1TimeStart := time.Now()
	if config.Config.DataSetSize > config.Config.DataSetBreaker {
		count := config.Config.DataSetSize / config.Config.DataSetBreaker
		var wg sync.WaitGroup
		for i := 0; i < count; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				startTime := time.Now()
				t1 := (config.Config.DataSetBreaker * config.Config.Type1) / 100
				t2 := (config.Config.DataSetBreaker * config.Config.Type2) / 100
				t3 := (config.Config.DataSetBreaker * config.Config.Type3) / 100
				t4 := (config.Config.DataSetBreaker * config.Config.Type4) / 100
				data := generateItemDataWithCountNewImpl(t1, t2, t3, t4)
				endTimeGenerate := time.Now()
				log.Printf("generateItemDataWithCountNewImpl() took %s time", time.Since(startTime).String())
				content, err := gocsv.MarshalString(data)
				log.Printf("gocsv.MarshalString() took %s time", time.Since(endTimeGenerate).String())
				if err != nil {
					log.Fatalf("gocsv.MarshalString() err: %v", err)
				}
				endTimeGocsv := time.Now()
				GenerateFileWithExtAndSuffix([]byte(content), fmt.Sprint("_", i+1), "csv")
				log.Printf("GenerateFileWithExtAndSuffix took %s time", time.Since(endTimeGocsv).String())
			}(i)

		}
		wg.Wait()
	} else {
		GenerateCSVData()
	}
	log.Printf("GenerateCSVDataV1() took %s time", time.Since(generateCSVDataV1TimeStart).String())
	return nil
}

func GenerateCSVDataV2MemConserve() error {
	generateCSVDataV1TimeStart := time.Now()
	log.Printf("GenerateCSVDataV2MemConserve started at %s", time.Now().String())

	if config.Config.DataSetSize > config.Config.DataSetBreaker {
		count := config.Config.DataSetSize / config.Config.DataSetBreaker

		for i := 0; i < count; i++ {

			startTime := time.Now()
			t1 := (config.Config.DataSetBreaker * config.Config.Type1) / 100
			t2 := (config.Config.DataSetBreaker * config.Config.Type2) / 100
			t3 := (config.Config.DataSetBreaker * config.Config.Type3) / 100
			t4 := (config.Config.DataSetBreaker * config.Config.Type4) / 100
			data := generateItemDataWithCountNewImpl(t1, t2, t3, t4)
			endTimeGenerate := time.Now()
			log.Printf("generateItemDataWithCountNewImpl() took %s time", time.Since(startTime).String())
			content, err := gocsv.MarshalString(data)
			log.Printf("gocsv.MarshalString() took %s time", time.Since(endTimeGenerate).String())
			if err != nil {
				log.Fatalf("gocsv.MarshalString() err: %v", err)
			}
			endTimeGocsv := time.Now()
			GenerateFileWithExtAndSuffix([]byte(content), fmt.Sprint("_", i+1), "csv")
			log.Printf("GenerateFileWithExtAndSuffix took %s time", time.Since(endTimeGocsv).String())

		}

	} else {
		GenerateCSVData()
	}
	log.Printf("GenerateCSVDataV2MemConserve ended at %s", time.Now().String())
	log.Printf("GenerateCSVDataV2MemConserve() took %s time", time.Since(generateCSVDataV1TimeStart).String())
	return nil
}

func GenerateCSVData() error {
	startTime := time.Now()
	data := generateItemData()
	endTimeGenerate := time.Now()
	log.Printf("generateItemData() took %s time", time.Since(startTime).String())
	content, err := gocsv.MarshalString(data)
	log.Printf("gocsv.MarshalString() took %s time", time.Since(endTimeGenerate).String())
	if err != nil {
		log.Fatalf("gocsv.MarshalString() err: %v", err)
	}
	endTimeGocsv := time.Now()
	GenerateFileWithExt([]byte(content), "csv")
	log.Printf("GenerateFileWithExt took %s time", time.Since(endTimeGocsv).String())
	return nil
}

// GenerateFile will generated file with extension
func GenerateFileWithExt(b []byte, extension string) {
	file, err := os.Create(fmt.Sprint(config.Config.OutputFileName, ".", extension))
	if err != nil {
		log.Fatalf("error while Generating file, %v", err)
	}

	n, err := file.Write(b)
	log.Printf("wrote %d bytes to file ", n)
	if err != nil {
		log.Fatalf("error while Generating file, %v", err)
	}
}

// GenerateFile will generated file with extension
func GenerateFileWithExtAndSuffix(b []byte, suffix, extension string) {
	dirName := "outputData"
	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	file, err := os.Create(fmt.Sprint(dirName, "/", config.Config.OutputFileName, suffix, ".", extension))
	if err != nil {
		log.Fatalf("error while Generating file, %v", err)
	}

	n, err := file.Write(b)
	log.Printf("wrote %d bytes to file ", n)
	if err != nil {
		log.Fatalf("error while Generating file, %v", err)
	}
}

// GenerateFile will generated file
func GenerateFile(b []byte) {
	file, err := os.Create(fmt.Sprint(config.Config.OutputFileName, ".", config.Config.OutputFormat))
	if err != nil {
		log.Fatalf("error while Generating file, %v", err)
	}

	n, err := file.Write(b)
	log.Printf("wrote %d bytes to file ", n)
	if err != nil {
		log.Fatalf("error while Generating file, %v", err)
	}
}
