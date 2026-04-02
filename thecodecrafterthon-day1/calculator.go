package main

import (
	"fmt"
)

func main() {

	var FirstNumber float64
	var operator string
	var SecondNumber float64
	var help string
	var quit string
start:

	for {
		fmt.Println("Enter FirstNumber:")
		fmt.Scanln(&FirstNumber)
		fmt.Println("operators:+ - * /")
		fmt.Scanln(&operator)
		fmt.Println("Enter SecondNumber:")
		fmt.Scan(&SecondNumber)

		
		switch operator {
		case "+":
			
		case "-":
			
		case "*":
			
		case "/":

		default:
			fmt.Println("operator not found:ivalid command, use help")
			fmt.Scanln(&help)
			fmt.Println(`
	(operators for calculation)
+ :the addition sign is for adding
- :the minus sign is for subtraction
* :the multiplication sign  is for multiplying
/ :the division sign is for dividing`)
			//return
		}
		break	
	}

	if operator == "+"{
		result:= FirstNumber + SecondNumber
		fmt.Println()
		fmt.Println("Result:", FirstNumber, "+", SecondNumber, "=", result)
	}
	if operator == "-"{
		result:= FirstNumber - SecondNumber
		fmt.Println()
		fmt.Println("Result:", FirstNumber, "-", SecondNumber, "=", result)

	}
	if operator == "*"{
		result:= FirstNumber * SecondNumber
		fmt.Println()
		fmt.Println("Result:", FirstNumber, "*", SecondNumber, "=", result)
	}
	if operator == "/"{
				if SecondNumber == 0{
			fmt.Println("invalid")
			
		}
		result:= FirstNumber / SecondNumber
		fmt.Println()
		fmt.Println("Result:", FirstNumber, "/", SecondNumber, "=", result)
	}
	fmt.Println("Do you want to quit?")
		fmt.Scanln(&quit)

		switch quit{
		case "No":
			goto start
		case "yes":
			fmt.Println("good bye....")
			return
		}
}
