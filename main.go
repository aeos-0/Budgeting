// Add saving to databases???
/*This can be improved by upddating total variable everytime addBill is called instead
	of using a for loop when adding up values*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processNum() float32 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ")
		panic(err)
	}
	//Fix input problem removing input spaces
	input = strings.TrimSpace(input)
	//Convert string to float
	inputNum, err := strconv.ParseFloat(input, 32)
	if err != nil {
		//Float values in the string will return an error
		fmt.Println("Error parsing input:", err)
	}
	return float32(inputNum)
}

func processChar() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ")
		panic(err)
	}
	//Weird input fix
	runes := []rune(input)
	input = string(runes[0])
	return input
}

func processString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ")
		panic(err)
	}
	return input
}

func addBill(bills *map[string]float32) {
	fmt.Println("Now you will need to add your monthly costs")

	//Start loop while !done
	done := false
	for !done {
		fmt.Println("Enter the name of the bill you would like to add")
		name := processString()
		fmt.Println("Now enter the amount it costs")
		cost := processNum()

		//First check if value exists and warn of overwrite
		_, ok := (*bills)[name]
		if ok {
			fmt.Println("A value has already been assigned to the key and will be overwritten")
			(*bills)[name] = cost
			fmt.Print(name)
			fmt.Println(" has been updated!")
		} else {
			(*bills)[name] = cost
			fmt.Print(name)
			fmt.Println(" has been added!")
		}

		
		
		fmt.Println("Would you like to add another invoice? Please type 'y' or 'n.'")
		fmt.Println("Invalid values will exit the program")

		answer := processChar()
		if answer == "y" || answer == "Y" {
			continue
		} else if answer == "n" || answer == "N" {
			done = true
		} else {
			panic("Invalid input")
		}
	}
}

func main() {
	bills := make(map[string]float32)

	fmt.Println("Welcome to the budgeting program!")
	fmt.Println("We can first begin by seleting your yearly budget in usd: ")

	//Collect budget from user input
	budget := processNum()
	
	addBill(&bills)

	//Start the main for loop
	/*From here you have many options
	Add more bills
	Remove bills
	Calculate month invoice
	How much remaining money per month and per year
	Change budget
	Quit program
	*/
	//done := false
	for {
		fmt.Println("What would you like to do now? Here are your options: ")
		fmt.Println("'a' - Add more invoices")
		fmt.Println("'d' - Remove invoices")
		fmt.Println("'m' - Calculate monthly invoice")
		fmt.Println("'?' - Check how much money you have to spare")
		fmt.Println("'c' - Change your yearly budget")
		fmt.Println("'q' - Exit the program")

		input := processChar()
		//Go provides automatic break statements which differ from C++
		switch input {
		case "a":
			//Add more invoices
			addBill(&bills)
		case "d":
			//Remove invoices
			fmt.Println("What bill would you like to remove:")
			for key := range bills {
				fmt.Println(key)
			}
			bill := processString()
			_, ok := bills[bill]
			if !ok {
				fmt.Println("This value doesn't exist")
				break
			}
			delete(bills, bill)
			fmt.Println(bill, "has been successfully removed")
		case "m":
			//Calculate monthly invoices
			var total float32
			for _, value := range bills {
				total += value
			}
			fmt.Println("Your monthly invoice is: ", total)
		case "?":
			//Check left over money monthly and yearly
			var total float32
			for _, value := range bills {
				total += value
			}
			fmt.Println("Per month you have ", budget / 12 - total, " to spare")
			fmt.Println("Per year you have $", budget - (total * 12), "to spare")
		case "c":
			//Change your yearly budget
			fmt.Println("What would you like to change the budget to")
			budget = processNum()
			fmt.Print("Your yearly budget is now: ")
			fmt.Println(budget)
		case "q":
			//Quit
			fmt.Println("Thank you for using our program!")
			os.Exit(0)
		default:
			fmt.Println("Invalid input, please try again")
		}
	}
}
