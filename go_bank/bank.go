package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

var filename = "balance.txt"
func main() {

	accountBalance:= fileops.ReadFloatValue(filename)
	fmt.Println("Welcome to go bank")
	fmt.Printf("======\nOur headquarters is %s %s\n", randomdata.City(), randomdata.Country(0))
	fmt.Printf("You can contact us 27/7 on %s\n======\n", randomdata.PhoneNumber())

	choice(accountBalance)
	fmt.Println("Thanks for choosing Go bank!")

}


