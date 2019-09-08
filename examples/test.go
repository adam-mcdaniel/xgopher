package main

import (
	"fmt"
	"math"

	. "github.com/adam-mcdaniel/xmachine-go"
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
		xasm.Push(NewString("n"))
		xasm.Store()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("m"))
			xasm.Store()
			xasm.Push(NewString("m"))
			xasm.Load()
			xasm.Copy()
			xasm.Push(NewString("n"))
			xasm.Load()
			xasm.Copy()
			xasm.Push(NewString("mul"))
			xasm.Load()
			xasm.Call()
		}, xasm.Duplicate()))
	}, xasm.Duplicate()))
	xasm.Copy()
	xasm.Push(NewString("multiply"))
	xasm.Store()
	xasm.Push(NewNumber(2))
	xasm.Copy()
	xasm.Push(NewString("multiply"))
	xasm.Load()
	xasm.Call()
	xasm.Copy()
	xasm.Push(NewString("double"))
	xasm.Store()
	xasm.Push(NewNumber(3))
	xasm.Copy()
	xasm.Push(NewString("multiply"))
	xasm.Load()
	xasm.Call()
	xasm.Copy()
	xasm.Push(NewString("triple"))
	xasm.Store()
	xasm.Push(NewFunction(func(xasm *Machine) {
		xasm.Push(NewString("a"))
		xasm.Store()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("b"))
			xasm.Store()
			xasm.Push(NewString("a"))
			xasm.Load()
		}, xasm.Duplicate()))
	}, xasm.Duplicate()))
	xasm.Copy()
	xasm.Push(NewString("True"))
	xasm.Store()
	xasm.Push(NewFunction(func(xasm *Machine) {
		xasm.Push(NewString("a"))
		xasm.Store()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("b"))
			xasm.Store()
			xasm.Push(NewString("b"))
			xasm.Load()
		}, xasm.Duplicate()))
	}, xasm.Duplicate()))
	xasm.Copy()
	xasm.Push(NewString("False"))
	xasm.Store()
	xasm.Push(NewFunction(func(xasm *Machine) {
		xasm.Push(NewString("c"))
		xasm.Store()
		xasm.Push(NewFunction(func(xasm *Machine) {
			xasm.Push(NewString("a"))
			xasm.Store()
			xasm.Push(NewFunction(func(xasm *Machine) {
				xasm.Push(NewString("b"))
				xasm.Store()
				xasm.Push(NewString("b"))
				xasm.Load()
				xasm.Copy()
				xasm.Push(NewString("a"))
				xasm.Load()
				xasm.Copy()
				xasm.Push(NewString("c"))
				xasm.Load()
				xasm.Call()
				xasm.Call()
			}, xasm.Duplicate()))
		}, xasm.Duplicate()))
	}, xasm.Duplicate()))
	xasm.Copy()
	xasm.Push(NewString("If"))
	xasm.Store()
	xasm.Push(NewNumber(2))
	xasm.Copy()
	xasm.Push(NewNumber(1))
	xasm.Copy()
	xasm.Push(NewString("False"))
	xasm.Load()
	xasm.Copy()
	xasm.Push(NewString("If"))
	xasm.Load()
	xasm.Call()
	xasm.Call()
	xasm.Call()
	xasm.Copy()
	xasm.Push(NewString("println"))
	xasm.Load()
	xasm.Call()
	xasm.Push(NewNumber(5))
	xasm.Copy()
	xasm.Push(NewString("double"))
	xasm.Load()
	xasm.Call()
	xasm.Copy()
	xasm.Push(NewString("println"))
	xasm.Load()
	xasm.Call()
	xasm.Push(NewNumber(5))
	xasm.Copy()
	xasm.Push(NewString("triple"))
	xasm.Load()
	xasm.Call()
	xasm.Copy()
	xasm.Push(NewString("println"))
	xasm.Load()
	xasm.Call()

	fmt.Println(xasm)
}

// fn rem(xasm: &mut Machine) {
//     let m = xasm.get_arg::<f64>();
//     let n = xasm.get_arg::<f64>();

//     xasm.push(
//         Value::number(m % n)
//     );
// }

// fn main() {
//     let mut xasm = Machine::new();
//     xasm.push(Value::function(dict, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("dict"));
//     xasm.store();

//     xasm.push(Value::function(list, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("list"));
//     xasm.store();
//     xasm.push(Value::function(len, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("len"));
//     xasm.store();
//     xasm.push(Value::function(push, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("push"));
//     xasm.store();
//     xasm.push(Value::function(pop, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("pop"));
//     xasm.store();

//     xasm.push(Value::function(print, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("print"));
//     xasm.store();
//     xasm.push(Value::function(println, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("println"));
//     xasm.store();
//     xasm.push(Value::function(new, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("new"));
//     xasm.store();

//     xasm.push(Value::function(add, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("add"));
//     xasm.store();
//     xasm.push(Value::function(sub, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("sub"));
//     xasm.store();
//     xasm.push(Value::function(mul, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("mul"));
//     xasm.store();
//     xasm.push(Value::function(div, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("div"));
//     xasm.store();
//     xasm.push(Value::function(rem, &xasm));
//     xasm.copy();
//     xasm.push(Value::string("rem"));
//     xasm.store();
//  xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("object"));
// 	xasm.store();
// 	xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("object"));
// 	xasm.load();
// 	xasm.copy();
// 	xasm.push(Value::string("println"));
// 	xasm.load();
// 	xasm.call();
// 	}, &xasm));
// 	xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("object"));
// 	xasm.load();
// 	xasm.push(Value::string("to_str"));
// 	xasm.method_call();
// 	xasm.copy();
// 	xasm.push(Value::string("println"));
// 	xasm.load();
// 	xasm.call();
// 	}, &xasm));
// 	xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("object"));
// 	xasm.load();
// 	xasm.push(Value::string("to_str"));
// 	xasm.index();
// 	}, &xasm));
// 	xasm.if_then_else();
// 	}, &xasm));
// 	xasm.copy();
// 	xasm.push(Value::string("print"));
// 	xasm.store();
// 	xasm.push(Value::string("string"));
// 	xasm.copy();
// 	xasm.push(Value::string("print"));
// 	xasm.load();
// 	xasm.call();
// 	xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("dict"));
// 	xasm.load();
// 	xasm.call();
// 	xasm.copy();
// 	xasm.push(Value::string("self"));
// 	xasm.store();
// 	xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("self"));
// 	xasm.store();
// 	xasm.push(Value::string("self"));
// 	xasm.load();
// 	}, &xasm));
// 	xasm.copy();
// 	xasm.push(Value::string("self"));
// 	xasm.load();
// 	xasm.push(Value::string("new"));
// 	xasm.index();
// 	xasm.assign();
// 	xasm.push(Value::function(|xasm: &mut Machine| {xasm.push(Value::string("self"));
// 	xasm.store();
// 	xasm.push(Value::string("TESTING!!!"));
// 	}, &xasm));
// 	xasm.copy();
// 	xasm.push(Value::string("self"));
// 	xasm.load();
// 	xasm.push(Value::string("to_str"));
// 	xasm.index();
// 	xasm.assign();
// 	xasm.push(Value::string("self"));
// 	xasm.load();
// 	}, &xasm));
// 	xasm.copy();
// 	xasm.push(Value::string("Test"));
// 	xasm.store();
// 	xasm.push(Value::string("Test"));
// 	xasm.load();
// 	xasm.copy();
// 	xasm.push(Value::string("new"));
// 	xasm.load();
// 	xasm.call();
// 	xasm.copy();
// 	xasm.push(Value::string("print"));
// 	xasm.load();
// 	xasm.call();

// }
