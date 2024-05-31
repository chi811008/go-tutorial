package main

import "fmt"

// go version 1.23+
func rangeNew() {
	// int
	fmt.Println("int i := 0; i < 10; i ++")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("range 10")
	for i := range 10 {
		fmt.Println(i)
	}
	fmt.Println("go1.22 has lift-off!")

	// function, 0 values f func(func()bool)
	for v := range seq1 {
		fmt.Println("in loop, v:", v)
	}
}

func seq1(yield func(int) bool) {
	fmt.Println("in seq1")
}
