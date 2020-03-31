package command

import (
	"kitabisa-test/receiver"
)

type SumCommand struct {
	Calculator receiver.ICalculator
}

func (c *SumCommand) Execute() {
	c.Calculator.Sum()
}
