package main

import "fmt"

func main() {
	var accountBalance = 1000.00

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

		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much would you like to deposit ? ")
			var depositAmount float64

			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount entered, must be greater then 0:", depositAmount)
				continue
			}

			accountBalance += depositAmount
			fmt.Println("The account balance have been updated ! New Amount is:", accountBalance)
		} else if choice == 3 {
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
		} else if choice == 4 {
			fmt.Println("Exited your account successfully !")
			break
		} else {
			fmt.Println("Wrong choice selected !")
			break
		}
	}

	fmt.Println("Thank for choosing our bank !")

}
