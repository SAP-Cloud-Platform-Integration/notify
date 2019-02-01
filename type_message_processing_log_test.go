package main

import (
	"io/ioutil"
	"testing"

	"gotest.tools/assert"
)

func TestParseMessageProcessingLogFromString(t *testing.T) {

	dat, _ := ioutil.ReadFile("./type_message_processing_log_test.json")

	m := ParseMessageProcessingLogFromString(dat)

	assert.DeepEqual(t, *m.D.Count, "102")

}
