package main

import "fmt"

//function
func main() {
	fmt.Println("Hello, World!")
	//initiallize the variables
	var x int = 5
	var y int = 10
	swap(x, y)

}

//add function
func swap(a int, b int) {
	fmt.Println(a, b)
}

/*
Concurrency
handle multiple things at a same time*/
