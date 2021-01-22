package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1 + 1) //2
	fmt.Println("1 + 1 =", 1.0 + 1.0) //2
	fmt.Println(len("Hello, World")) //12
	fmt.Println("Hello, World"[1]) //101
	fmt.Println("Hello," + "World") //Hello, World
	fmt.Println(true && true) //true
	fmt.Println(true && false) //false
	fmt.Println(true || true) //true
	fmt.Println(true || false) //true
	fmt.Println(!true) //false

	fmt.Println(32132 * 42452) //1364067664
}


