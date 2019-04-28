package controller

import (
	"eRecipet/terminal/iohandler"
	"eRecipet/terminal/model"
)

/*
This is demo program for eRecipt system
The target is terminal for displaying QR code

MessageCenter
This object contol message sending

Author	:Kazufumi Watanabe
Create	:Apr-28,2019
Updates
*/

/*
MessageCenter :Interface of message center which handles message delivery
*/
type MessageCenter interface {
	ReadCVS(path string) (*model.Receipt, bool)
	CreateMessage(recipet *model.Receipt) (*model.Message, bool)
}

// message center
type messageCenter struct {
}

/*
NewMessageCenter :initializer
*/
func NewMessageCenter() MessageCenter {
	obj := new(messageCenter)

	return obj
}

/*
ReadCVS
	in	:path string
	out	:bool
*/
func (m messageCenter) ReadCVS(path string) (*model.Receipt, bool) {
	// check path
	if path == "" {
		return nil, false
	}

	// initializer I/O handler
	reader, err := iohandler.NewCSVReader(path)
	if err != nil {
		return nil, false
	}

	// get read data
	buffer := reader.Buffer()
	if buffer == nil {
		return nil, false
	}

	// initialize parser
	parser := NewCSVParser(buffer)
	if parser == nil {
		return nil, false
	}

	// get recipet data
	recipet := parser.ParseRecipt()
	if recipet == nil {
		return nil, false
	}

	return recipet, true
}

/*
CreateMessage
	in	:*recipet *model.Receipt
	out	:*model.Message, bool
*/
func (m *messageCenter) CreateMessage(recipet *model.Receipt) (*model.Message, bool) {
	if recipet == nil {
		return nil, false
	}

	message := model.NewMessage(recipet)
	if message == nil {
		return nil, false
	}

	return message, true
}
