package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

func (s *Student) DoHomeWork(number int) {
	fmt.Printf("%s is doing homework %d\n", s.name, number)
}

func main4() {
	s := Student{name: "heli"}
	v := reflect.ValueOf(&s)
	methodV := v.MethodByName("DoHomeWork")
	if methodV.IsValid() {
		in := []reflect.Value{reflect.ValueOf(55)}
		methodV.Call(in)
	}
}