package main

import (
	"fmt"
	"math"

	. "github.com/adam-mcdaniel/xgopher"
)

func dict(m *Machine) {
	m.Push(NewEmptyTree())
}

func list(m *Machine) {
	m.Push(NewEmptyList())
}

func push(m *Machine) {
	list := m.Pop().Slice()
	item := m.Pop()

	m.Push(NewList(append(list, item)))
}

func pop(m *Machine) {
	list := m.Pop().Slice()

	item, list := list[len(list)-1], list[:len(list)-1]
	m.Push(item)
	m.Push(NewList(list))
}

func length(m *Machine) {
	m.Push(NewNumber(float64(len(m.Pop().Slice()))))
}

func print(m *Machine) {
	fmt.Print(*m.Pop())
}

func println(m *Machine) {
	fmt.Println(*m.Pop())
}

func new(m *Machine) {
	m.Call()
	m.Push(NewString("new"))
	m.MethodCall()
}

func add(m *Machine) {
	a := m.Pop()
	b := m.Pop()

	m.Push(NewNumber(a.Number() + b.Number()))
}

func sub(m *Machine) {
	a := m.Pop()
	b := m.Pop()

	m.Push(NewNumber(a.Number() - b.Number()))
}

func mul(m *Machine) {
	a := m.Pop()
	b := m.Pop()

	m.Push(NewNumber(a.Number() * b.Number()))
}

func div(m *Machine) {
	a := m.Pop()
	b := m.Pop()

	m.Push(NewNumber(a.Number() / b.Number()))
}

func rem(m *Machine) {
	a := m.Pop()
	b := m.Pop()

	m.Push(NewNumber(math.Mod(a.Number(), b.Number())))
}

func main() {
	xasm := MakeMachine()
	xasm.Push(NewFunction(dict, xasm))
	xasm.Push(NewString("dict"))
	xasm.Store()
	xasm.Push(NewFunction(list, xasm))
	xasm.Push(NewString("list"))
	xasm.Store()
	xasm.Push(NewFunction(length, xasm))
	xasm.Push(NewString("len"))
	xasm.Store()
	xasm.Push(NewFunction(push, xasm))
	xasm.Push(NewString("push"))
	xasm.Store()
	xasm.Push(NewFunction(pop, xasm))
	xasm.Push(NewString("pop"))
	xasm.Store()
	xasm.Push(NewFunction(print, xasm))
	xasm.Push(NewString("print"))
	xasm.Store()
	xasm.Push(NewFunction(println, xasm))
	xasm.Push(NewString("println"))
	xasm.Store()
	xasm.Push(NewFunction(new, xasm))
	xasm.Push(NewString("new"))
	xasm.Store()
	xasm.Push(NewFunction(add, xasm))
	xasm.Push(NewString("add"))
	xasm.Store()
	xasm.Push(NewFunction(sub, xasm))
	xasm.Push(NewString("sub"))
	xasm.Store()
	xasm.Push(NewFunction(mul, xasm))
	xasm.Push(NewString("mul"))
	xasm.Store()
	xasm.Push(NewFunction(div, xasm))
	xasm.Push(NewString("div"))
	xasm.Store()
	xasm.Push(NewFunction(rem, xasm))
	xasm.Push(NewString("rem"))
	xasm.Store()

	xasm.Push(NewFunction(func(xasm *Machine) {
		xasm.Push(NewString("dict"))
		xasm.Load()
		xasm.Call()
		xasm.Copy()
		xasm.Push(NewString("self"))
		xasm.Store()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("self"))
			xasm.Store()
			xasm.Push(NewString("list"))
			xasm.Load()
			xasm.Call()
			xasm.Copy()
			xasm.Push(NewString("self"))
			xasm.Load()
			xasm.Push(NewString("list"))
			xasm.Index()
			xasm.Assign()
			xasm.Push(NewString("self"))
			xasm.Load()
		}, xasm.Duplicate()))
		xasm.Copy()
		xasm.Push(NewString("self"))
		xasm.Load()
		xasm.Push(NewString("new"))
		xasm.Index()
		xasm.Assign()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("self"))
			xasm.Store()
			xasm.Push(NewFunction(func(xasm *Machine) {}, xasm.Duplicate()))
			xasm.Push(NewFunction(func(xasm *Machine) {
				xasm.Push(NewString("self"))
				xasm.Load()
				xasm.Push(NewString("list"))
				xasm.Index()
				xasm.Copy()
				xasm.Push(NewString("pop"))
				xasm.Load()
				xasm.Call()
				xasm.Copy()
				xasm.Push(NewString("self"))
				xasm.Load()
				xasm.Push(NewString("list"))
				xasm.Index()
				xasm.Assign()
			}, xasm.Duplicate()))
			xasm.Push(NewFunction(func(xasm *Machine) {
				xasm.Push(NewString("self"))
				xasm.Load()
				xasm.Push(NewString("list"))
				xasm.Index()
				xasm.Copy()
				xasm.Push(NewString("len"))
				xasm.Load()
				xasm.Call()
			}, xasm.Duplicate()))
			xasm.IfThenElse()
		}, xasm.Duplicate()))
		xasm.Copy()
		xasm.Push(NewString("self"))
		xasm.Load()
		xasm.Push(NewString("pop"))
		xasm.Index()
		xasm.Assign()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("self"))
			xasm.Store()
			xasm.Push(NewString("value"))
			xasm.Store()
			xasm.Push(NewString("value"))
			xasm.Load()
			xasm.Copy()
			xasm.Push(NewString("self"))
			xasm.Load()
			xasm.Push(NewString("list"))
			xasm.Index()
			xasm.Copy()
			xasm.Push(NewString("push"))
			xasm.Load()
			xasm.Call()
			xasm.Copy()
			xasm.Push(NewString("self"))
			xasm.Load()
			xasm.Push(NewString("list"))
			xasm.Index()
			xasm.Assign()
		}, xasm.Duplicate()))
		xasm.Copy()
		xasm.Push(NewString("self"))
		xasm.Load()
		xasm.Push(NewString("push"))
		xasm.Index()
		xasm.Assign()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("self"))
			xasm.Store()
			xasm.Push(NewString("self"))
			xasm.Load()
			xasm.Push(NewString("list"))
			xasm.Index()
			xasm.Copy()
			xasm.Push(NewString("len"))
			xasm.Load()
			xasm.Call()
		}, xasm.Duplicate()))
		xasm.Copy()
		xasm.Push(NewString("self"))
		xasm.Load()
		xasm.Push(NewString("len"))
		xasm.Index()
		xasm.Assign()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("self"))
			xasm.Store()
			xasm.Push(NewString("n"))
			xasm.Store()
			xasm.Push(NewString("self"))
			xasm.Load()
			xasm.Push(NewString("list"))
			xasm.Index()
			xasm.Push(NewString("n"))
			xasm.Load()
			xasm.Index()
		}, xasm.Duplicate()))
		xasm.Copy()
		xasm.Push(NewString("self"))
		xasm.Load()
		xasm.Push(NewString("index"))
		xasm.Index()
		xasm.Assign()
		xasm.Push(NewString("self"))
		xasm.Load()
	}, xasm.Duplicate()))
	xasm.Copy()
	xasm.Push(NewString("List"))
	xasm.Store()
	xasm.Push(NewString("List"))
	xasm.Load()
	xasm.Copy()
	xasm.Push(NewString("new"))
	xasm.Load()
	xasm.Call()
	xasm.Copy()
	xasm.Push(NewString("l"))
	xasm.Store()
	xasm.Push(NewNumber(10))
	xasm.Copy()
	xasm.Push(NewString("n"))
	xasm.Store()
	xasm.Push(NewFunction(func(xasm *Machine) {
		xasm.Push(NewString("n"))
		xasm.Load()
		xasm.Copy()
		xasm.Push(NewString("l"))
		xasm.Load()
		xasm.Push(NewString("push"))
		xasm.MethodCall()
		xasm.Push(NewNumber(1))
		xasm.Copy()
		xasm.Push(NewString("n"))
		xasm.Load()
		xasm.Copy()
		xasm.Push(NewString("sub"))
		xasm.Load()
		xasm.Call()
		xasm.Copy()
		xasm.Push(NewString("n"))
		xasm.Store()
	}, xasm.Duplicate()))
	xasm.Push(NewFunction(func(xasm *Machine) {
		xasm.Push(NewString("n"))
		xasm.Load()
	}, xasm.Duplicate()))
	xasm.WhileLoop()
	xasm.Push(NewString("l"))
	xasm.Load()
	xasm.Copy()
	xasm.Push(NewString("println"))
	xasm.Load()
	xasm.Call()

}
