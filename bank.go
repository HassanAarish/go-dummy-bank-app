package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0664)
}

func main() {
	var accountBalance = getBalanceFromFile()

	for {
		fmt.Println("Welcome to go bank !")
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Printf("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
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
			writeBalanceToFile(accountBalance)

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
			writeBalanceToFile(accountBalance)

		case 4:
			fmt.Println("Exited your account successfully !")
			return

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
