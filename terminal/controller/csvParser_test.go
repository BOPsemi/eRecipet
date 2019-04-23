package controller

import (
	"eRecipet/terminal/iohandler"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	CSVFILEPATH = "/Users/kazufumiwatanabe/go/src/eRecipet/terminal/data/recipet.csv"
)

func MocData() [][]string {
	reader, err := iohandler.NewCSVReader(CSVFILEPATH)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return reader.Buffer()
}

func TestNewCSVParser(t *testing.T) {
	data := MocData()
	parser := NewCSVParser(data)
	recipet := parser.ParseRecipt()

	assert.NotNil(t, recipet)
	fmt.Println(recipet.HiddenInfo.IDCode)
}
