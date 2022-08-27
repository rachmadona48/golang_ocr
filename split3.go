package main

import (
	"fmt"
	"strings"
)

func main() {

	// myString := "Hello! This is a golangcode.com test ;)"
	// strParts := strings.Split(myString, "!")

	myString := "Management Fee                              1  PRS        361,748                  361,749"
	strmodel := strings.Split(myString, "                              ")
	model := strmodel[0]
	strqty := strings.Split(strmodel[1], "        ")
	qty := strqty[0]
	price_idr := strqty[1]
	stramount := strings.Split(strmodel[1], "                  ")
	amount := stramount[1]
	fmt.Println(stramount[0], "-", stramount[1])
	fmt.Println(model, "|", qty, "|", price_idr, "|", amount)
}
