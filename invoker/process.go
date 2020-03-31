package invoker

import (
	"kitabisa-test/command"
)

type Process struct {
	Command command.ICommand
}

func (p *Process) process() {
	p.Command.Execute()
}
