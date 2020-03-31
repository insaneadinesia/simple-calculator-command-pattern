package main

import (
	"bufio"
	"fmt"
	"kitabisa-test/command"
	"kitabisa-test/receiver"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("KITABISA BACKEND TEST")
	fmt.Println("=====================")
	fmt.Println("Choose the following simple program")
	fmt.Println("1. Sum of 2 numbers")
	fmt.Println("2. Multiply of 2 numbers")
	fmt.Println("3. Print first N prime number")
	fmt.Println("4. Print first N fibonacci")
	fmt.Println()

	fmt.Print("Enter your choice (1, 2, 3, 4) : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimRight(choice, "\r\n")

	switch choice {
	case "1":
		input := GetInput("sum")
		if input == nil {
			fmt.Println("Invalid input. Close.")
			break
		}

		sumCommand := &command.SumCommand{Calculator: input}
		sumCommand.Calculator.Sum()
		break
	case "2":
		input := GetInput("multiply")
		if input == nil {
			fmt.Println("Invalid input. Close.")
			break
		}

		multiplyCommand := &command.MultiplyCommand{Calculator: input}
		multiplyCommand.Calculator.Multiply()
		break
	case "3":
		input := GetInput("prime")
		if input == nil {
			fmt.Println("Invalid input. Close.")
			break
		}

		primeCommand := &command.PrimeCommand{Calculator: input}
		primeCommand.Calculator.Prime()
		break
	case "4":
		input := GetInput("fibonacci")
		if input == nil {
			fmt.Println("Invalid input. Close.")
			break
		}

		fibonacciCommand := &command.FibonacciCommand{Calculator: input}
		fibonacciCommand.Calculator.Fibonacci()
		break
	default:
		fmt.Println("Invalid input. Close.")
	}
}

func GetInput(inputType string) *receiver.Calculator {
	reader := bufio.NewReader(os.Stdin)

	x := 0
	y := 0

	switch inputType {
	case "sum", "multiply":
		fmt.Print("Enter the [x, y] (example : 3, 2) : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		input = strings.Replace(input, " ", "", -1)

		inputs := strings.Split(input, ",")

		if len(inputs) != 2 {
			return nil
		}

		x, _ = strconv.Atoi(inputs[0])
		y, _ = strconv.Atoi(inputs[1])
		break

	case "prime", "fibonacci":
		fmt.Print("Enter N number to generate : ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimRight(input1, "\r\n")

		x, _ = strconv.Atoi(input1)
		break
	}

	calculator := &receiver.Calculator{
		X: x,
		Y: y,
	}

	return calculator
}
