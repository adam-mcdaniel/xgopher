package xgopher

import "fmt"

type Error struct {
	s string
}

func (e Error) Error() string {
	return e.s
}

func MakeError(s string) Error {
	return Error{s}
}

type Machine struct {
	stack     []*Value
	registers map[string]*Value
}

func (a Machine) eq(b Machine) bool {
	for i := range a.stack {
		if i >= len(b.stack) {
			return false
		}
		if !(a.stack[i].Eq(b.stack[i])).Bool() {
			return false
		}
	}

	for key, first := range a.registers {
		if second, ok := b.registers[key]; ok {
			if !(first.Eq(second)).Bool() {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func MakeMachine() Machine {
	return Machine{stack: []*Value{}, registers: make(map[string]*Value)}
}

func (m *Machine) Push(value *Value) {
	m.stack = append(m.stack, value)
}

func (m *Machine) Pop() *Value {
	result := NewError("Popped from empty stack, called function with too few arguments")
	if len(m.stack) == 0 {
		return result
	}

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
	if val, ok := m.registers[key.String()]; ok {
		m.Push(val)
	} else {
		m.Push(NewError(fmt.Sprintf("No register named %v", key.String())))
	}
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
	getCondition := func(m *Machine) bool {
		return m.Pop().Bool()
	}

	condition.callGlobal(m)
	for getCondition(m) {
		body.callGlobal(m)
		condition.callGlobal(m)
	}
}

func (m *Machine) IfThenElse() {
	condition := m.Pop()
	thenFn := m.Pop()
	elseFn := m.Pop()
	getCondition := func(m *Machine) bool {
		return m.Pop().Bool()
	}

	condition.callGlobal(m)
	if getCondition(m) {
		thenFn.callGlobal(m)
	} else {
		elseFn.callGlobal(m)
	}
}

func (m Machine) Duplicate() Machine {
	newMachine := MakeMachine()
	for _, item := range m.stack {
		newMachine.Push(item.Copy().Copy())
	}

	for key, value := range m.registers {
		newMachine.registers[key] = value.Copy().Copy()
	}

	return newMachine
}

func (m Machine) String() string {
	result := "Machine {\n\tstack: ["
	for i, value := range m.stack {
		result += fmt.Sprintf("%v", value)
		if i != len(m.stack)-1 {
			result += ", "
		}
	}
	result += "]\n\theap:  {"
	singleElement := true
	for i, value := range m.registers {
		result += fmt.Sprintf("\"%v\": %v", i, value)
		result += ", "
		singleElement = false
	}

	if !singleElement {
		result = result[:len(result)-2]
	}

	result += "}\n}"

	return result
}
