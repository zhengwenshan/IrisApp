package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := new(Test)
	of := reflect.TypeOf(t)
	method, ok := of.MethodByName("A")
	if ok {
		inputs := make([]reflect.Value, method.Type.NumIn())
		inputs[0] = reflect.ValueOf(t)
		method.Func.Call(inputs)
	}
}

type Test struct {
}

func (t *Test) A() {
	fmt.Print("AAAAA")
}
