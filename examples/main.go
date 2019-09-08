package main

import (
	"fmt"

	. "github.com/adam-mcdaniel/xmachine-go"
)

func main() {
	m := MakeMachine()
	m.Push(NewFunction(func(m *Machine) {
		fmt.Println("FUNCTION CALL!")
		m.Push(NewString("RETURN VALUE"))
	}, m))
	m.Push(NewString("testing"))
	m.Store()
	m.Push(NewString("testing"))
	m.Load()
	m.Call()

	m.Push(NewNumber(5))
	m.Copy()
	m.Push(NewString("testing"))
	m.Load()
	m.Assign()
	fmt.Println(m)
}
