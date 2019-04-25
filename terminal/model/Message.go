package model

import (
	"github.com/google/uuid"
)

/*
This is demo program for eReceipt system
The target is terminal for displaying QR code

Message object is a object for sending message to server side

Author	:Kazufumi Watanabe
Create	:Apr-24,2019
Updates
*/

/*
Message :objct of message
*/
type Message struct {
	recepietData *Receipt // body of message structure
	uuid         string   // object serial ID
	hashKey      string   // hashKey for verification check
}

// UUID generator
func uuidGen() string {
	uuid, _ := uuid.NewUUID()
	return uuid.String()
}

// hashKeyGen :hash key generator
func hashKeyGen(data *Receipt) string {
	if data == nil {
		return ""
	}

	//timeStamp := time.Now()

	return ""
}

/*
NewMessage :initializer of Message model
*/
func NewMessage(data *Receipt) *Message {
	if data == nil {
		return nil
	}

	// return new object
	return &Message{
		recepietData: data,
		uuid:         uuidGen(),
		hashKey:      "hige",
	}
}
