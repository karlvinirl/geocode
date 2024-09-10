package main

import "fmt"
import "example.com/bank/fileops"

func choice (accountBalance float64) {
	var close = false
	for close != true {
	
		presentOptions()
		var choice int
		fmt.Println("Choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Your balance is %v\n", fileops.ReadFloatValue(filename))
		case 2:
			var deposit float64
			fmt.Println("Amount of deposit: ")
			fmt.Scan(&deposit)
			accountBalance += deposit
			fileops.WriteFloatValue(filename, accountBalance)
			fmt.Printf("Your new Balance is: %v\n", accountBalance)
		case 3:
			fmt.Printf("How much would you like to withdraw ? You available balance is %v\n", fileops.ReadFloatValue(filename))
			var withdrawal float64
			fmt.Scan(&withdrawal)
			if accountBalance-withdrawal < 0 {
				fmt.Println("Withdrawal exceeds account funds, please try again or choose another option")
			} else {
				accountBalance -= withdrawal
				fileops.WriteFloatValue(filename, accountBalance)
				fmt.Printf("Please take your money, Your new balance is %0.2f\n", accountBalance)
			}
		case 4:
			fmt.Println("Goodybe!")
			close = true
		}
		
	}
}