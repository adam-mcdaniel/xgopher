package xgopher

import (
	"fmt"
	"strconv"
)

type Type int

const (
	StringType   = iota
	NumberType   = iota
	ListType     = iota
	FunctionType = iota
	TreeType     = iota
	ErrorType    = iota
	NoneType     = iota
)

type Value struct {
	valueType Type
	str       string
	num       float64
	list      []*Value
	fn        Function
	tree      map[string]*Value
	err       string
}

func NewNone() *Value {
	return &Value{valueType: NoneType}
}

func NewString(str string) *Value {
	return &Value{valueType: StringType, str: str}
}

func NewNumber(num float64) *Value {
	return &Value{valueType: NumberType, num: num}
}

func NewEmptyList() *Value {
	return NewList([]*Value{})
}

func NewList(list []*Value) *Value {
	return &Value{valueType: ListType, list: list}
}

func NewFunction(Function func(*Machine), m Machine) *Value {
	return &Value{valueType: FunctionType, fn: makeFunction(Function, m)}
}

func NewEmptyTree() *Value {
	return NewTree(make(map[string]*Value))
}

func NewTree(tree map[string]*Value) *Value {
	return &Value{valueType: TreeType, tree: tree}
}

func NewError(err string) *Value {
	return &Value{valueType: ErrorType, err: err}
}

func (v *Value) Index(index string) *Value {
	switch v.valueType {
	case ListType:
		num, err := strconv.Atoi(index)
		if err != nil {
			return NewNone()
		}

		for num >= len(v.list) {
			v.list = append(v.list, NewNone())
		}

		return v.list[num]
	case StringType:
		num, err := strconv.Atoi(index)
		if err != nil {
			return NewNone()
		}

		return NewString(string(v.str[num]))
	case TreeType:
		if val, ok := v.tree[index]; ok {
			return val
		} else {
			v.tree[index] = NewNone()
			return v.tree[index]
		}
	default:
		return NewNone()
	}
}

func (v Value) Bool() bool {
	switch v.valueType {
	case StringType:
		return len(v.str) > 0
	case NumberType:
		return v.num != 0
	case ListType:
		return len(v.list) > 0
	case FunctionType:
		return true
	case TreeType:
		return len(v.tree) > 0
	case ErrorType:
		return false
	default:
		return false
	}
}

func (v Value) String() string {
	switch v.valueType {
	case StringType:
		return v.str
	case NumberType:
		return fmt.Sprintf("%v", v.num)
	case ListType:
		return fmt.Sprintf("%v", v.list)
	case FunctionType:
		return fmt.Sprintf("%v", v.fn)
	case TreeType:
		result := "{"
		singleElement := true
		for i, value := range v.tree {
			result += fmt.Sprintf("%v: %v", i, value.String())
			result += ", "
			singleElement = false
		}

		if !singleElement {
			result = result[:len(result)-2]
		}

		result += "}"
		return result
	case ErrorType:
		return fmt.Sprintf("<Exception: '%v'>", v.err)
	default:
		return "None"
	}
}

func (v Value) Slice() []*Value {
	return v.list
}

func (v Value) Tree() map[string]*Value {
	return v.tree
}

func (v Value) Number() float64 {
	return v.num
}

func (v Value) Copy() *Value {
	switch v.valueType {
	case StringType:
		return NewString(v.str)
	case NumberType:
		return NewNumber(v.num)
	case ListType:
		newList := []*Value{}
		for _, value := range v.list {
			newList = append(newList, value.Copy())
		}
		return NewList(newList)
	case FunctionType:
		return NewFunction(v.fn.fn, v.fn.context)
	case TreeType:
		newMap := make(map[string]*Value)
		for key, value := range v.tree {
			newMap[key] = value.Copy()
		}
		return NewTree(newMap)
	case ErrorType:
		return NewError(v.err)
	default:
		return NewNone()
	}
}

func (v Value) call(machine *Machine) {
	if v.valueType == FunctionType {
		tempMachine := v.fn.getContext()
		tempMachine.stack = machine.stack
		v.fn.call(&tempMachine)
		machine.stack = tempMachine.stack
	}
}

func (v Value) callGlobal(machine *Machine) {
	if v.valueType == FunctionType {
		v.fn.call(machine)
	}
}
