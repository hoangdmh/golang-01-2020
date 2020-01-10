package main

import "fmt"

/*
func func_name(params type) return type {
	//code here
}
*/

func hello() {
	fmt.Println("Hello Go")
}

func sayHello(name string) {
	fmt.Println("Say hello", name)
}

func getName(name string) string {
	result := fmt.Sprintf("say hello %s", name)
	return result
}

//multiple return values
func recInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

// named return values
func recInfo1(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	fmt.Println("===========> Function in Go <===========")
	sayHello("Hoang")

	fmt.Println(getName("Nam"))

	// w, h, area := recInfo(10, 20)
	// fmt.Println("Width is ", w)
	// fmt.Println("Height is", h)
	// fmt.Println("Dien tich ", area)

	w, h, isSquare := recInfo1(20, 20)
	if isSquare {
		fmt.Println("this is square")
	} else {
		fmt.Println("Width is ", w)
		fmt.Println("Height is", h)
	}

	fmt.Println("===========> End Function in Go <===========")
}
