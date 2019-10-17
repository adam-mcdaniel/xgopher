package xgopher

import "fmt"

type Function struct {
	fn      func(*Machine)
	context map[string]*Value
}

func makeFunction(fn func(*Machine), context Machine) Function {
	return Function{fn, context.registers}
}

func (fn Function) call(machine *Machine) {
	(fn.fn)(machine)
}

func (fn Function) getContext() Machine {
	machine := MakeMachine()
	machine.registers = fn.context
	return machine
}

func (fn Function) String() string {
	return fmt.Sprintf("<fn at %p>", fn.fn)
}
