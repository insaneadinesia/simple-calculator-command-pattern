package command

import (
	"kitabisa-test/receiver"
)

type PrimeCommand struct {
	Calculator receiver.ICalculator
}

func (c *PrimeCommand) Execute() {
	c.Calculator.Prime()
}
