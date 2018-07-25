package process

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
)

// CHANGEME - modify function param type to fit your event type used in handler
type EventProcessor interface {
	ProcessRecord(record events.Record) error
}

// CHANGEME - Add fields in this struct for contextual type objects you want to make available to the ProcessRecord call.
type eventProcessorImpl struct {
	logger                 zerolog.Logger
	someContextualParam    string
}

// CHANGEME - Define a struct here that matches your specific event structure
type myEvent struct {
	MyEventField     string `json:"webhookType"`
}

// CHANGEME - modify the param to use the event type member based on the event type used in handler
func (p eventProcessorImpl) ProcessRecord(dataRecord events.Record) error {
	dataBytes := dataRecord.Data

	//Parse data into Json
	parsedJSON := myEvent{}
	err := json.Unmarshal(dataBytes, &parsedJSON)
	if err != nil {
		p.logger.Error().Err(err).Msg("Failure to parse Event Data JSON")
		return err
	}
	// CHANGEME - Put business logic here to process your event. Do any validation checks
	// and processing for your event here.

	return nil
}

func New(logger zerolog.Logger, param string) EventProcessor {
	return eventProcessorImpl{
		logger:            logger,
		someContextualParam: param,
	}
}
