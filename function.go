package xgopher

import "fmt"

type Function struct {
	fn      func(*Machine)
	context Machine
}

func makeFunction(fn func(*Machine), context Machine) Function {
	return Function{fn, context}
}

func (fn Function) call(machine *Machine) {
	(fn.fn)(machine)
}

func (fn Function) getContext() Machine {
	return fn.context
}

func (fn Function) String() string {
	return fmt.Sprintf("<fn at %p>", fn.fn)
}
