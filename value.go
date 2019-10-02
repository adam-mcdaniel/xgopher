package xgopher

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func NewBool(value bool) *Value {
	if value {
		return NewNumber(1)
	} else {
		return NewNumber(0)
	}
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

func (v Value) Type() Type {
	return v.valueType
}

func (v *Value) Index(index string) *Value {
	switch v.valueType {
	case ListType:
		num, err := strconv.Atoi(index)
		if err != nil {
			return NewError("Can't index list with non-integer")
		}

		for num >= len(v.list) {
			v.list = append(v.list, NewNone())
		}

		return v.list[num]
	case StringType:
		num, err := strconv.Atoi(index)
		if err != nil {
			return NewError("Can't index string with non-integer")
		}

		if len(v.str) > num && num >= 0 {
			return NewString(string(v.str[num]))
		} else {
			return NewError("String index out of bounds")
		}

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
		return math.Abs(v.num) > 0.0000000001
	case ListType:
		return len(v.list) > 0
	case FunctionType:
		return true
	case TreeType:
		return len(v.tree) > 0
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
		result := "["
		singleElement := true
		for _, value := range v.list {
			result += fmt.Sprintf("%v", value.String())
			result += ", "
			singleElement = false
		}

		if !singleElement {
			result = result[:len(result)-2]
		}

		result += "]"
		return result
	case FunctionType:
		return fmt.Sprintf("%v", v.fn)
	case TreeType:
		result := "{"
		singleElement := true
		for i, value := range v.tree {
			result += fmt.Sprintf("\"%v\": %v", i, value.String())
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

func (v Value) IsNone() bool {
	return v.valueType == NoneType
}

func (v Value) Str() string {
	return v.str
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

func (v Value) IntoIter() []*Value {
	switch v.valueType {
	case ListType:
		return v.Slice()
	case StringType:
		result := []*Value{}
		for ch := range v.str {
			result = append(result, NewString(string(ch)))
		}
		return result
	default:
		return []*Value{}
	}
}


func (a Value) Add(b *Value) *Value {
	aType := a.valueType
	bType := b.valueType
	if aType == StringType && bType == StringType {
		return NewString(a.str + b.str)
	} else if aType == NumberType && bType == NumberType {
		return NewNumber(a.num + b.num)
	} else if aType == ListType && bType == ListType {
		return NewList(append(a.list, b.list...))
	}
	return NewError(fmt.Sprintf("Could not add %v and %v", a, b))
}

func (a Value) Sub(b *Value) *Value {
	aType := a.valueType
	bType := b.valueType
	if aType == NumberType && bType == NumberType {
		return NewNumber(a.num - b.num)
	}
	return NewError(fmt.Sprintf("Could not subtract %v and %v", a, b))
}

func (a Value) Mul(b *Value) *Value {
	aType := a.valueType
	bType := b.valueType
	if aType == NumberType && bType == NumberType {
		return NewNumber(a.num * b.num)
	} else if aType == StringType && bType == NumberType {
		return NewString(strings.Repeat(a.str, int(b.num)))
	}
	return NewError(fmt.Sprintf("Could not multiply %v and %v", a, b))
}

func (a Value) Div(b *Value) *Value {
	aType := a.valueType
	bType := b.valueType
	if aType == NumberType && bType == NumberType {
		return NewNumber(a.num / b.num)
	}
	return NewError(fmt.Sprintf("Could not divide %v and %v", a, b))
}

func (a Value) Rem(b *Value) *Value {
	aType := a.valueType
	bType := b.valueType
	if aType == NumberType && bType == NumberType {
		return NewNumber(math.Mod(a.num, b.num))
	}
	return NewError(fmt.Sprintf("Could not find the remainder of %v and %v", a, b))
}

func (a Value) Not() *Value {
	if a.valueType == NumberType {
		if int(a.num) == 0 {
			return NewNumber(1)
		} else {
			return NewNumber(0)
		}
	}
	return NewError(fmt.Sprintf("Could not negate %v", a))
}

func (a Value) Eq(b *Value) *Value {
	aType := a.valueType
	bType := b.valueType
	if aType != bType {
		return NewBool(false)
	}

	if aType == NumberType && bType == NumberType {
		return NewBool(a.num == b.num)
	} else if aType == StringType && bType == StringType {
		return NewBool(a.str == b.str)
	} else if aType == ListType && bType == ListType {
		for i := range a.list {
			if !(a.list[i].Eq(b.list[i])).Bool() {
				return NewBool(false)
			}
		}
		return NewBool(true)
	} else if aType == FunctionType && bType == FunctionType {
		return NewBool(a.String() == b.String())
	} else if aType == TreeType && bType == TreeType {
		for key, first := range a.tree {
			if second, ok := b.tree[key]; ok {
				if !(first.Eq(second)).Bool() {
					return NewBool(false)
				}
			} else {
				return NewBool(false)
			}
		}
		return NewBool(true)
	} else if aType == ErrorType && bType == ErrorType {
		return NewBool(a.err == b.err)
	} else if aType == NoneType && bType == NoneType {
		return NewBool(true)
	}

	return NewBool(false)
}
