package main

import "fmt"

func mapwithnum() {

	var NumFromMap = map[int]string{

		1: "One",
		2: "Two",
	}

	ValFromMap := NumFromMap[1]

	fmt.Println(ValFromMap)
	fmt.Println(NumFromMap[2])
}
