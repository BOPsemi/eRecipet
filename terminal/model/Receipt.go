package model

/*
This is demo program for eReceipt system
The target is terminal for displaying QR code

Receipt object is a model for Receipt data

Author	:Kazufumi Watanabe
Create	:Apr-21,2019
Updates
*/

/*
ItemInfo :the defintion of the purchased item information
Name 	:Item name
Number	:Purchased item number
Price	:Purchased each item price
*/
type ItemInfo struct {
	Name   string
	Number int
	Price  float64
}

/*
HiddenInfo :describe hidden info in Receipt
*/
type HiddenInfo struct {
	PurchasedMethod int
	Provider        string
	IDCode          int
}

/*
Receipt :Receipt info
*/
type Receipt struct {
	Date       string
	Time       string
	TerminalID string
	Tax        int
	ItemInfo   []ItemInfo
	HiddenInfo *HiddenInfo
}
