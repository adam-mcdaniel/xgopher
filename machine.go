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

func MakeMachine() Machine {
	return Machine{stack: []*Value{}, registers: make(map[string]*Value)}
}

func (m *Machine) Push(value *Value) {
	m.stack = append(m.stack, value)
}

func (m *Machine) Pop() (*Value, error) {
	result := NewNone()
	if len(m.stack) == 0 {
		return result, MakeError("No value on stack")
	}
	result, m.stack = m.stack[len(m.stack)-1], m.stack[:len(m.stack)-1]

	return result, nil
}

func (m *Machine) Call() {
	if value, err := m.Pop(); err == nil {
		value.call(m)
	}
}

func (m *Machine) Copy() {
	if value, err := m.Pop(); err == nil {
		m.Push(value.Copy())
	}
}

func (m *Machine) Store() {
	if key, err := m.Pop(); err == nil {
		if value, err := m.Pop(); err == nil {
			m.registers[key.String()] = value
		}
	}
}

func (m *Machine) Load() {
	key, err := m.Pop()
	if val, ok := m.registers[key.String()]; ok && (err == nil) {
		m.Push(val)
	} else {
		m.Push(NewError(fmt.Sprintf("No register named %v", key.String())))
	}
}

func (m *Machine) Assign() {
	if reference, err := m.Pop(); err == nil {
		if value, err := m.Pop(); err == nil {
			*reference = *value
		}
	}
}

func (m *Machine) Index() {
	if index, err := m.Pop(); err == nil {
		if table, err := m.Pop(); err == nil {
			m.Push(table.Index(index.String()))
		}
	}
}

func (m *Machine) MethodCall() {
	if index, err := m.Pop(); err == nil {
		if table, err := m.Pop(); err == nil {
			m.Push(table)
			m.Push(table)
			m.Push(index)
			m.Index()
			m.Call()
		}
	}
}

func (m *Machine) WhileLoop() {
	if condition, err := m.Pop(); err == nil {
		if body, err := m.Pop(); err == nil {
			getCondition := func(m *Machine) bool {
				result, err := m.Pop()
				if err != nil {
					return false
				}
				return result.Bool()
			}

			condition.callGlobal(m)
			for getCondition(m) {
				body.callGlobal(m)
				condition.callGlobal(m)
			}
		}
	}
}

func (m *Machine) IfThenElse() {
	if condition, err := m.Pop(); err == nil {
		if thenFn, err := m.Pop(); err == nil {
			if elseFn, err := m.Pop(); err == nil {
				getCondition := func(m *Machine) bool {
					result, err := m.Pop()
					if err != nil {
						return false
					}
					return result.Bool()
				}

				condition.callGlobal(m)
				if getCondition(m) {
					thenFn.callGlobal(m)
				} else {
					elseFn.callGlobal(m)
				}
			}
		}
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
