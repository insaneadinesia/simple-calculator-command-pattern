package main

import (
	"bufio"
	"fmt"
	"kitabisa-test/command"
	"kitabisa-test/receiver"
	"os"
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
		calculator := &receiver.Calculator{
			X: 2,
			Y: 4,
		}

		sumCommand := &command.SumCommand{Calculator: calculator}
		sumCommand.Calculator.Sum()
		break
	case "2":
		calculator := &receiver.Calculator{
			X: 2,
			Y: 4,
		}

		multiplyCommand := &command.MultiplyCommand{Calculator: calculator}
		multiplyCommand.Calculator.Multiply()
		break
	case "3":
		calculator := &receiver.Calculator{
			X: 10,
		}

		primeCommand := &command.PrimeCommand{Calculator: calculator}
		primeCommand.Calculator.Prime()
		break
	case "4":
		fmt.Println("Fibonacci")
		break
	default:
		fmt.Println("Invalid input. Close.")
	}
}
