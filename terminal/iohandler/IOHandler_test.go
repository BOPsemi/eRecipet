package iohandler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	CSVFILEPATH = "/Users/kazufumiwatanabe/go/src/eRecipet/terminal/data/recipet.csv"
)

func TestNewIOHandler(t *testing.T) {
	assert.NotNil(t, NewIOHandler(CSVFILEPATH))
}

func TestNewCSVReader(t *testing.T) {
	reader, err := NewCSVReader(CSVFILEPATH)
	if err != nil {
		t.Fail()
	}

	assert.NotNil(t, reader)
}

func TestBuffer(t *testing.T) {
	reader, err := NewCSVReader(CSVFILEPATH)
	if err != nil {
		t.Fail()
	}

	buffer := reader.Buffer()
	assert.NotNil(t, buffer)
	fmt.Println(buffer)
}
