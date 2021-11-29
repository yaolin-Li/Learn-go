package main

import (
	"fmt"
	"reflect"
)

func Mymain3() {
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	fmt.Printf("is float canset: %v canAddr : %v\n",v.CanSet(), v.CanAddr())
	VP := reflect.ValueOf(&fVar)
	fmt.Printf("is float canset: %v canAddr : %v\n",VP.Elem().CanSet(), VP.Elem().CanAddr())
	VP.Elem().SetFloat(2.333333)
	fmt.Println(fVar)
}
