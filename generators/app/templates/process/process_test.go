package process

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// CHANGEME Put any mocks you need for your logic testing here

func readJsonFromFile(t *testing.T, inputFile string) []byte {
	inputJson, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJson
}

// CHANGEME - sample test code. Modify types and logic to fit your process.go code.
func Test_Handler_Process_DataRecord_Happy_Path(t *testing.T) {
	asserter := assert.New(t)
	buf := new(bytes.Buffer)
	logger := zerolog.New(buf)

	// 1. read JSON from file
	inputJson := readJsonFromFile(t, "../testdata/event-happy-path.json")
	var Event events.Event
	if err := json.Unmarshal(inputJson, &Event); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	var EventRecord events.EventRecord
	if err := json.Unmarshal(inputJson, &EventRecord); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// Test your event specific logic here

	eventProcessor := New(logger, "some param")
	err := eventProcessor.ProcessRecord(EventRecord.Event)

	asserter.Nil(err)

	// assert any other conditions needed to be checked
}

func Test_Handler_Process_DataRecord_Negative_Case(t *testing.T) {
	asserter := assert.New(t)
	buf := new(bytes.Buffer)
	logger := zerolog.New(buf)

	// 1. read JSON from file
	inputJson := readJsonFromFile(t, "../testdata/event-negative-case.json")
	var Event events.Event
	if err := json.Unmarshal(inputJson, &Event); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	var EventRecord events.EventRecord
	if err := json.Unmarshal(inputJson, &EventRecord); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	eventProcessor := New(logger, "some param")
	err := eventProcessor.ProcessRecord(EventRecord.Event)

	asserter.Nil(err)

	// assert any other conditions that should be checked
}


