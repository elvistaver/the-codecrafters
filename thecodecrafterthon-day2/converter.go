package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hex(hext string) int64 {

	hexacon, err := strconv.ParseInt(hext, 16, 64)
	if err != nil {
		fmt.Println("❌invalid hex input.Try again...")
	}
	return hexacon
}

func bin(bina string) int64 {
	bindec, err := strconv.ParseInt(bina, 2, 64)
	if err != nil {
		fmt.Println("❌invalid binary input.Try again...")
	}
	return bindec
}

func dec(deci string) (string, string) {
	d, err := strconv.ParseInt(deci, 10, 64)
	if err != nil {
		fmt.Println("❌invalid decimal input.Try again...")
	}
	decima := strconv.FormatInt(d, 2)
	decim := strconv.FormatInt(d, 16)
	return decima, strings.ToUpper(decim)

}

func main() {

	var end string
	var input string
	var conv string

	for {
		fmt.Println("Base Converter")
		fmt.Println("_____________________________________________")
	restart:
	start:
		fmt.Println("converter Base List:")
		fmt.Println("1. Hex-base 16")
		fmt.Println("2. Bin-base 2")
		fmt.Println("3. Dec-10")
		fmt.Println("_____________________________________________")
		fmt.Println("Choose a converter Base")
		fmt.Scanln(&conv)
		fmt.Println("_____________________________________________")
		fmt.Println("Input converter base number type")
		fmt.Scanln(&input)
		fmt.Println("_____________________________________________")

		switch conv {
		case "1":
			fmt.Println("Result:", "=", hex(input))
			fmt.Println()
		case "2":
			fmt.Println("Result:", "=", bin(input))
			fmt.Println()
		case "3":
			bin1, hex2 := dec(input)
			fmt.Println("Binary:", bin1)
			fmt.Println("Hex:", hex2)
		default:
			fmt.Println("❌invalid option: choose a valid convert base")
			fmt.Println("_____________________________________________")
			goto restart
		}
		fmt.Println("Done✅")
		fmt.Println("_____________________________________________")
	begin:
		fmt.Println("do you wish you to continue?")
		fmt.Println("1.yes")
		fmt.Println("2.No")
		fmt.Println("_____________________________________________")
		fmt.Println("enter option")
		fmt.Scanln(&end)
		fmt.Println("_____________________________________________")

		switch end {
		case "1":

		case "2":
		case end:
			fmt.Println("choose a valid option")
			fmt.Println("_____________________________________________")
			goto begin
		}

		if end == "1" {
			goto start
		}
		if end == "2" {
			fmt.Println("exiting...")
			fmt.Println("program sucessfully exited✅")
			return
		}
		break
	}
}
