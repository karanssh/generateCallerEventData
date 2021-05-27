package eventgen

import (
	"generateEventData/config"
	"generateEventData/datadefinition"
	"generateEventData/datagen"

	"github.com/google/uuid"
)

// generateEvent generated a random event
func GenerateEvent(eventType int) datadefinition.Event {
	var event datadefinition.Event
	event.CallingNumber = datagen.RandomNumberInRange(config.Config.PhoneNoLow, config.Config.PhoneNoHigh)
	event.CalledNumber = datagen.RandomNumberInRange(config.Config.PhoneNoLow, config.Config.PhoneNoHigh)
	event.DurationSeconds = datagen.RandomNumberInRange(config.Config.DurationLow, config.Config.DurationHigh)
	event.EventRef = uuid.New().String()
	event.EventSource = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Location = datagen.RandomStringFromRunes(config.Config.RandomTextSize) //TODO: add cities list
	event.EventType = eventType
	event.EventDate = datagen.RandomTime()
	event.Attr1 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr2 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr3 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr4 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr5 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr6 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr7 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)
	event.Attr8 = datagen.RandomStringFromRunes(config.Config.RandomTextSize)

	return event
}

// GenerateEvents generates events with a given event type
func GenerateEvents(numberOfEvents int, eventType int) []datadefinition.Event {
	output := make([]datadefinition.Event, numberOfEvents)
	for i := 0; i < numberOfEvents; i++ {
		output[i] = GenerateEvent(eventType)
	}
	return output
}
