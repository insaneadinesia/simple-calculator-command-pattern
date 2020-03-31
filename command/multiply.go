package command

import (
	"kitabisa-test/receiver"
)

type MultiplyCommand struct {
	Calculator receiver.ICalculator
}

func (c *MultiplyCommand) Execute() {
	c.Calculator.Multiply()
}
