package controller

import (
	"eRecipet/terminal/model"
	"strconv"
	"strings"
)

/*
This is demo program for eRecipt system
The target is terminal for displaying QR code

CSV Parser object description
This object analyzes CSV file, and return data object

Author	:Kazufumi Watanabe
Create	:Apr-18,2019
Updates
*/

/*
CSVParser : parser for CSV file
*/
type CSVParser struct {
	rawData [][]string
}

/*
NewCSVParser :initializer
*/
func NewCSVParser(data [][]string) *CSVParser {
	// check in data size and null
	if data == nil {
		return nil
	}

	// initialize raw data
	obj := new(CSVParser)
	obj.rawData = data

	return obj
}

// extractTheFirstElem
func (c *CSVParser) extractTheFirstElem(line []string) string {
	if len(line) == 0 && line == nil {
		return ""
	}

	return line[0]
}

/*
ParseRecipt :parse recipt data
*/
func (c *CSVParser) ParseRecipt() *model.Receipt {
	if c.rawData == nil {
		return nil
	}

	// create new recipet object as data model
	recipetData := new(model.Receipt)
	if recipetData == nil {
		return nil
	}

	// index info
	dataStartIndex := 0
	hiddenDataStartIndex := 0

	// readline
	for index, strline := range c.rawData {
		switch index {
		case 0:
			// created data
			recipetData.Date = c.extractTheFirstElem(strline)

		case 1:
			// created time
			recipetData.Time = c.extractTheFirstElem(strline)

		case 2:
			// terminal ID
			recipetData.TerminalID = c.extractTheFirstElem(strline)

		case 3:
			// tax
			recipetData.Tax, _ = strconv.Atoi(c.extractTheFirstElem(strline))

		default:
			if strline[0] == "rrr" {
				dataStartIndex = index + 1
			}

			if strline[0] == "hhh" {
				hiddenDataStartIndex = index + 1
			}
		}
	}

	// initialize stocker
	recipetData.ItemInfo = make([]model.ItemInfo, 0)

	// data body update section
	// --- etract purchased data ---
	purchasedData := c.rawData[dataStartIndex : hiddenDataStartIndex-1]
	for _, dataInfo := range purchasedData {

		// string separation by ","
		strs := strings.Split(dataInfo[0], ",")

		// initialize ItemInfo structure
		name := strs[0]
		number, _ := strconv.Atoi(strs[1])
		price, _ := strconv.ParseFloat(strs[2], 64)

		itemInfo := &model.ItemInfo{
			Name:   name,
			Number: number,
			Price:  price,
		}

		// data stacking
		recipetData.ItemInfo = append(recipetData.ItemInfo, *itemInfo)
	}

	// --- extract hidden data ---
	hiddenData := c.rawData[hiddenDataStartIndex:]

	method, _ := strconv.Atoi(c.extractTheFirstElem(hiddenData[0]))
	provider := c.extractTheFirstElem(hiddenData[1])
	idcode, _ := strconv.Atoi(c.extractTheFirstElem(hiddenData[2]))

	hiddenInfo := &model.HiddenInfo{
		PurchasedMethod: method,
		Provider:        provider,
		IDCode:          idcode,
	}
	recipetData.HiddenInfo = hiddenInfo

	return recipetData
}
