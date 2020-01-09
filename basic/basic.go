package main

import (
	"fmt"
)

/**
1. Khai báo biến
	var x int= 2
	x int, y int
		=> shortened x,y int

	// khac kieu dữ liệu
	var (
		name string
		age int
	)

	//shortened declare
	language := "Golang" // biến language có giá trị là Golang kiểu string


2. Types golang
	- bool -> true/false
	- string -> "hello"
	- int (so nguyên)
	- float (float32, float64)
	- byte
	- rune
	- complex64, complaex128 (số phức)


*/

func main() {
	// khai báo nhiều biến khác kiểu dữ liệu
	var (
		name string = "Hong"
		age  int    = 23
	)
	language := "Golang" //biến language có giá trị là Golang kiểu string

	fmt.Println("Basic language Golang")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(language)

	//rune

	var myString = "Tran Minh Tuấn"
	//runes := []rune(myString)
	for i := 0; i < len(myString); i++ {
		fmt.Printf("%c", myString[i])
	}
}
