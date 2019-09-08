package xmachine

import "fmt"

type Machine struct {
	stack     []*Value
	registers map[string]*Value
}

func MakeMachine() Machine {
	return Machine{stack: []*Value{}, registers: make(map[string]*Value)}
}

func (m *Machine) Push(value *Value) {
	m.stack = append(m.stack, value)
}

func (m *Machine) Pop() *Value {
	result := NewNone()
	result, m.stack = m.stack[len(m.stack)-1], m.stack[:len(m.stack)-1]
	return result
}

func (m *Machine) Call() {
	m.Pop().call(m)
}

func (m *Machine) Copy() {
	m.Push(m.Pop().Copy())
}

func (m *Machine) Store() {
	key := m.Pop()
	value := m.Pop()

	m.registers[key.String()] = value
}

func (m *Machine) Load() {
	key := m.Pop()

	m.Push(m.registers[key.String()])
}

func (m *Machine) Assign() {
	reference := m.Pop()
	value := m.Pop()
	*reference = *value
}

func (m *Machine) Index() {
	index := m.Pop()
	table := m.Pop()

	m.Push(table.Index(index.String()))
}

func (m *Machine) MethodCall() {
	index := m.Pop()
	table := m.Pop()

	m.Push(table)
	m.Push(table)
	m.Push(index)
	m.Index()
	m.Call()
}

func (m *Machine) WhileLoop() {
	condition := m.Pop()
	body := m.Pop()

	condition.callGlobal(m)
	for m.Pop().Bool() {
		body.callGlobal(m)
		condition.callGlobal(m)
	}
}

func (m *Machine) IfThenElse() {
	condition := m.Pop()
	thenFn := m.Pop()
	elseFn := m.Pop()

	condition.callGlobal(m)
	if m.Pop().Bool() {
		thenFn.callGlobal(m)
	} else {
		elseFn.callGlobal(m)
	}
}

func (m *Machine) Duplicate() Machine {
	newMachine := MakeMachine()
	for _, item := range m.stack {
		newMachine.Push(item.Copy())
	}

	for key, value := range m.registers {
		newMachine.registers[key] = value.Copy()
	}

	return newMachine
}

func (m Machine) String() string {
	result := "["
	for i, value := range m.stack {
		result += fmt.Sprintf("%v", value)
		if i != len(m.stack)-1 {
			result += ", "
		}
	}
	result += "]\n{"
	singleElement := true
	for i, value := range m.registers {
		result += fmt.Sprintf("%v: %v", i, value)
		result += ", "
		singleElement = false
	}

	if !singleElement {
		result = result[:len(result)-2]
	}

	result += "}"

	return result
}
