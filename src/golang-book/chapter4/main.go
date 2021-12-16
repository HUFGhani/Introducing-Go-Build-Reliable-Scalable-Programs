package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	/*
		this long hand how write for loop
	*/
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	/*
		this short hand how write a for loop
	*/
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	/*
		this is how a if statement looks
	*/
	if i%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	/*
		you can add if statements inside for loop
	*/

	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	/*
		switch case instead of if statements
	*/

	l := 4
	switch l {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

}

}
