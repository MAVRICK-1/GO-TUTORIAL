package main

import (
	"fmt"
	"strconv"
)

func main() {
	var command string

	fmt.Print("Enter a command: ")
	fmt.Scanln(&command)

	// for _, i := range command {
	// 	fmt.Println("Command is: ", string(i))
	// }

	num1, err1 := strconv.Atoi(string(command[0])) // convert string to int Atoi = ASCII to integer
	num2, err2 := strconv.Atoi(string(command[2])) // convert string to int Atoi = ASCII to integer

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers")
		return
	}

	switch command[1] {
	case '+':
		fmt.Println(num1 + num2)
	case '-':
		fmt.Println(num1 - num2)
	case '*':
		fmt.Println(num1 * num2)
	case '/':
		if num2 != 0 {
			fmt.Println(num1 / num2)
		} else {
			fmt.Println("Cannot divide by zero")
		}
	default:
		fmt.Println("Invalid command")
	}
}