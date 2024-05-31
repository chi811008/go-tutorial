package main

import (
	"fmt"
)

func rangeOri() {
	// slice
	fmt.Println("slice")
	numbers := []int{1, 2, 3, 4, 5}
	for idx, val := range numbers {
		fmt.Printf("index: %d, value: %d\n", idx, val)
	}
	// string
	fmt.Println("string")
	str := "Hello, Gamania!"
	for idx, val := range str {
		fmt.Printf("index: %d, value: %c\n", idx, val)
	}
	// array
	fmt.Println("array")
	arr := [5]int{1, 2, 3, 4, 5}
	for idx, val := range arr {
		fmt.Printf("index: %d, value: %d\n", idx, val)
	}
	// map
	fmt.Println("map")
	map1 := map[string]int{"遊戲": 1, "橘子": 2}
	for key, val := range map1 {
		fmt.Printf("key: %s, value: %d\n", key, val)
	}
	// channel
	fmt.Println("channel")
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for val := range ch {
		fmt.Printf("value: %d\n", val)
	}
}
