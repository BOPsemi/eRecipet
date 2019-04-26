package model

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	TIMESTAMP_LAYOUT = "2006-01-02 15:04:05" // layout for time stamp
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

// timeStamp :time stamp as string
func timeStamp() string {
	return time.Now().Format(TIMESTAMP_LAYOUT)
}

// objectToByte :convert structure to byte
func objectToByte(obj *Receipt) ([]byte, error) {
	if obj == nil {
		return nil, errors.New("Object data is empty")
	}

	// convert the inputted data to []byte
	byteData, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}

// hashKeyGen :hash key generator
func hashKeyGen(data *Receipt) (string, error) {
	if data == nil {
		return "", errors.New("object data is nil")
	}

	timeStampBytes := []byte(timeStamp())
	objectBytes, err := objectToByte(data)
	if err != nil {
		return "", err
	}

	// connect
	objectBytes = append(objectBytes, timeStampBytes...)

	// generate hash key from objectBytes
	hashKey := sha512.Sum512(objectBytes)
	hashKeyStr := hex.EncodeToString(hashKey[:])

	return hashKeyStr, nil
}

/*
NewMessage :initializer of Message model
*/
func NewMessage(data *Receipt) *Message {
	if data == nil {
		return nil
	}

	hashKeyStr, _ := hashKeyGen(data)

	// return new object
	return &Message{
		recepietData: data,
		uuid:         uuidGen(),
		hashKey:      hashKeyStr,
	}
}
