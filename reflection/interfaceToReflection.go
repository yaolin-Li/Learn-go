package main

import (
	"fmt"
	"reflect"
)

func printMeta(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	k := t.Kind()
	n := t.Name()
	fmt.Printf("Type: %s Kind: %s Name: %s Value:%v\n",t, k, n, v)
}

type handler func(int, int) int

func Mymain() {
	// pair: <value, type>

	var intVar int64 = 10
	stringVar := "hello"
	type book struct {
		name string
		pages int
	}
	sum := func(a, b int) int {
		return a + b
	}
	var sub handler = func(a, b int) int {
		return a - b
	}
	sli := make([]int, 5)

	printMeta(intVar)
	printMeta(stringVar)
	printMeta(book{
		name: "harry potter",
		pages: 500,
	})
	printMeta(sum)
	printMeta(sub)
	printMeta(sli)

}
