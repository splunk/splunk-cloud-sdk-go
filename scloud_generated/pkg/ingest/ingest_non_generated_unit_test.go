package ingest

import (
	"bufio"
	model "github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadBatch(t *testing.T) {
	s := "First event\n" +
		"Second event\n" +
		"Third event\n"
	reader := strings.NewReader(s)
	buf := bufio.NewReader(reader)

	output, err := readBatch(buf)
	require.Nil(t, err)
	require.NotNil(t, output)

	assert.Equal(t, 3, len(output))
}

func TestReadBatchEmptyFile(t *testing.T) {
	s := ""
	reader := strings.NewReader(s)
	buf := bufio.NewReader(reader)

	output, err := readBatch(buf)
	require.NotNil(t, err)
	assert.EqualError(t, err, "EOF")
	require.Nil(t, output)
}

func TestPostBatchRaw(t *testing.T) {
	batch := []string {"first event", "second event"}

	host := "hoststr"
	source := "sourcestr"
	sourcetype := "sourcetypestr"

	args := model.Event{
		Host:       &host,
		Source:     &source,
		Sourcetype: &sourcetype,
	}

	events, err := postBatchRaw(batch, args)
	require.Nil(t, err)

	assert.Equal(t, 2, len(events))
	assert.Equal(t, host, *events[0].Host)
	assert.Equal(t, source, *events[0].Source)
	assert.Equal(t, sourcetype, *events[0].Sourcetype)
}

func TestPostBatchJSON(t *testing.T) {
	batch := []string {"[\"first event\"]", "[\"second event\"]"}

	host := "hoststr"
	source := "sourcestr"
	sourcetype := "sourcetypestr"

	args := model.Event{
		Host:       &host,
		Source:     &source,
		Sourcetype: &sourcetype,
	}

	events, err := postBatchEventJSON(batch, args)
	require.Nil(t, err)

	assert.Equal(t, 2, len(events))
	assert.Equal(t, host, *events[0].Host)
	assert.Equal(t, source, *events[0].Source)
	assert.Equal(t, sourcetype, *events[0].Sourcetype)

}
