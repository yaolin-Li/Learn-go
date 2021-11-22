package main

import "fmt"

func main() {
	/*
	array := [5]string{"1", "2", "3", "4", "5"}
	array2 := [...]string{"1", "2"}
	fmt.Println(len(array))
	fmt.Println(cap(array))
	fmt.Println(len(array2))
	fmt.Println(cap(array2))

	array3 := array
	array3[0] = "another"
	fmt.Println(array3)
	fmt.Println(array)

	modify(array)
	fmt.Println(array)

	//切片 （起始的指针，长度，容量）
	slice := []string{"1", "3"}
	fmt.Println(len(slice), cap(slice))

	slice2 := make([]int, 2, 2)
	fmt.Println(slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 1)
	fmt.Println(slice2, len(slice2), cap(slice2))
	modify2(slice2)
	fmt.Println(slice2, len(slice2), cap(slice2))
	*/

	//映射map
	m := map[string]int{
		"ming":10,
		"zhang":13,
	}
	fmt.Println(m["ming"])
}


func modify2(slice []int)  {
	slice[0]=100
}

func modify(array [5]string) {
	array[0] = "123"
}
