package controller

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessageCenter(t *testing.T) {
	messageCenter := NewMessageCenter()
	assert.NotNil(t, messageCenter)
}

func TestReadCSV(t *testing.T) {

	messageCenter := NewMessageCenter()
	recipet, status := messageCenter.ReadCVS(CSVFILEPATH)
	assert.True(t, status)
	fmt.Println(recipet)

}

func TestCreateMessage(t *testing.T) {
	messageCenter := NewMessageCenter()
	recipet, status := messageCenter.ReadCVS(CSVFILEPATH)

	if status {
		message, status := messageCenter.CreateMessage(recipet)
		assert.True(t, status)
		fmt.Println(message)
	}

}
