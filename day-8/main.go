package main

import "fmt"

func main() {

	myLoop()
	myWhile()
	// myBreak()
	// myContinue()
	myArr()
}

func myLoop() {

	for i := 0; i < 5; i++ {
		fmt.Println("Hello World!!")
	}
}

func myWhile() {

	i := 1

	for i < 10 {
		fmt.Println("Hello using While Loop with For")
		i += 2
	}
}

func myBreak() {

	i := 0

	for {
		if i == 3 {
			break
		}
		fmt.Println(i)
		i++
	}

}

func myContinue() {

	for i := 0; i < 10; i++ {
		if i == 4 || i == 7 {
			continue
		}
		fmt.Println(i)
	}
}

func myArr() {

	arr := [4]int{5, 6, 7, 8}

	for index, value := range arr {
		fmt.Println(arr[index])
		fmt.Println(value)

	}
}
