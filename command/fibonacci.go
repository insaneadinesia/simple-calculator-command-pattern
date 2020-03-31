package command

import (
	"kitabisa-test/receiver"
)

type FibonacciCommand struct {
	Calculator receiver.ICalculator
}

func (c *FibonacciCommand) Execute() {
	c.Calculator.Fibonacci()
}
