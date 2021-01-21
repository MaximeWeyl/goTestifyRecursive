package behaviours

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
)

var (
	//X is a special argument that is replaced with the actual value being tested
	X specialArgument = x{}

	//F is a special argument that is replaced with the path of the field being tested
	//Use this for formatting your error message
	F specialArgument = f{}
)

type specialArgument interface {
	get(actualValueInterface interface{}, parentFieldName string) reflect.Value
}

type x struct{}

func (x x) get(actualValueInterface interface{}, _ string) reflect.Value {
	return reflect.ValueOf(actualValueInterface)
}

type f struct{}

func (f f) get(_ interface{}, parentFieldName string) reflect.Value {
	return reflect.ValueOf(fmt.Sprintf("{%s}", parentFieldName))
}

type behaviourFromAssert struct {
	functionValue reflect.Value
	args          []interface{}
}

//CheckField Performs the assertion with the underlying assert function and given arguments
//Special arguments (like X) are replaced before calling the assert function
func (b behaviourFromAssert) CheckField(t assert.TestingT, actualValueInterface interface{}, parentFieldName string) bool {
	functionValue := b.functionValue
	in := []reflect.Value{reflect.ValueOf(t)}
	for _, arg := range b.args {
		switch v := arg.(type) {
		case specialArgument:
			in = append(in, v.get(actualValueInterface, parentFieldName))
		default:
			in = append(in, reflect.ValueOf(arg))
		}
	}

	out := functionValue.Call(in)
	outBool := out[0].Interface().(bool)

	return outBool
}

//Func is a factory that returns a new behaviour from a testify assert function
func Func(function interface{}, args ...interface{}) FieldBehaviour {
	// Check type
	functionValue := reflect.ValueOf(function)
	if functionValue.Kind() != reflect.Func {
		panic("Must be called with a function")
	}
	funcType := functionValue.Type()

	// Check return value
	if numReturnValues := funcType.NumOut(); numReturnValues != 1 {
		panic(fmt.Sprintf("Must be called with a function with one return value. It had %d", numReturnValues))
	}
	if outKind := funcType.Out(0).Kind(); outKind != reflect.Bool {
		panic(fmt.Sprintf("Return value should be bool, but was %s", outKind))
	}

	// Check parameters
	numParameters := funcType.NumIn()
	if numParameters < 1 {
		panic(fmt.Sprintf("Return value should have at least 1 parameter, but had %d", numParameters))
	}

	return behaviourFromAssert{
		functionValue: functionValue,
		args:          args,
	}
}
