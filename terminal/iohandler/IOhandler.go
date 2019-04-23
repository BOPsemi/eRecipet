/*
This is demo program for eReceipt system
The target is terminal for displaying QR code

IOHandler object description, the object can handle IO and other interface R/W

Author	:Kazufumi Watanabe
Create	:Apr-16,2019
Updates
*/

package iohandler

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

/*
IOHandler :descritpion of IOHander object
*/
type IOHandler struct {
	err  error    // error
	file *os.File // pointer of file
}

/*
NewIOHandler :initializer of IOHandler
	in	:path string
	out	:*IOHandler
*/
func NewIOHandler(path string) *IOHandler {
	// create object
	obj := new(IOHandler)

	// check path
	if path == "" {
		return nil
	}

	// open file
	obj.file, obj.err = os.Open(path)
	if obj.err != nil {
		obj.file.Close()
		return nil
	}

	// return file
	return obj
}

/*
CSVReader :CSV reader object
*/
type CSVReader struct {
	ioHander *IOHandler // io hander object
	buffer   [][]string // buffer
}

/*
NewCSVReader :Initializer of CSV reader
	 in		:path string
	 out	:*CSVReader, error
*/
func NewCSVReader(path string) (*CSVReader, error) {
	// check file path
	if path == "" {
		return nil, errors.New("file path is empty")
	}

	// create object
	obj := new(CSVReader)
	if obj == nil {
		return nil, errors.New("CSV Reader initialization error")
	}

	// initializer file
	obj.ioHander = NewIOHandler(path)

	if obj.ioHander == nil {
		return nil, errors.New("File open error")
	}

	// read csv file
	reader := csv.NewReader(obj.ioHander.file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// --- debug ---
		// fmt.Println(line)
		// buffering
		obj.buffer = append(obj.buffer, line)
	}

	// close file
	obj.ioHander.file.Close()

	// return
	return obj, nil
}

/*
Buffer : getter of buffer
*/
func (c *CSVReader) Buffer() [][]string {
	return c.buffer
}
