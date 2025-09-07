package main

import (
	"fmt"

	"example.com/bank/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileOps.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("Error found: ", err)
		panic("Can't continue, sorry !")
	}

	fmt.Println("Welcome to go bank !")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		userOptions()

		var choice int
		fmt.Printf("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			fmt.Println("Exited your account successfully !")
			return

		case 1:
			fmt.Println("Your balance is:", accountBalance)

		case 2:
			fmt.Print("How much would you like to deposit ? ")
			var depositAmount float64

			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount entered, must be greater then 0:", depositAmount)
				continue
			}

			accountBalance += depositAmount
			fmt.Println("The account balance have been updated ! New Amount is:", accountBalance)
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Print("How much would you like to withdraw ? ")
			var withdrawAmount float64

			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount entered, must be greater then 0:", accountBalance)
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You can not withdraw more then you have:", accountBalance)
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("The account balance have been updated ! New Amount is:", accountBalance)
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 4:
			fmt.Println("Help Line Number: ", randomdata.PhoneNumber())
			continue

		default:
			fmt.Println("Wrong choice selected. Exitinng ...")
			return
			// break
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("How much would you like to deposit ? ")
		// 	var depositAmount float64

		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount entered, must be greater then 0:", depositAmount)
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("The account balance have been updated ! New Amount is:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("How much would you like to withdraw ? ")
		// 	var withdrawAmount float64

		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount entered, must be greater then 0:", accountBalance)
		// 		continue
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("You can not withdraw more then you have:", accountBalance)
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("The account balance have been updated ! New Amount is:", accountBalance)
		// } else if choice == 4 {
		// 	fmt.Println("Exited your account successfully !")
		// 	break
		// } else {
		// 	fmt.Println("Wrong choice selected !")
		// 	break
		// }
	}

	// fmt.Println("Thank for choosing our bank !")

}
