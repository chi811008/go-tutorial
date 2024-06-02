package main

import "fmt"

func main() {
	fmt.Println("normal for loop")
	ls1 := []int{6, 7, 8, 9, 10}
	for i, v := range ls1 {
		fmt.Println(i, v)
	}

	fmt.Println("basic yield for loop")
	for v := range seqYield {
		fmt.Println(v)
	}

	fmt.Println("yield for loop with number param")
	for v := range allInt(5) {
		fmt.Println(v)
	}

	fmt.Println("yield for loop with list param")
	for v := range allLsInt(ls1) {
		fmt.Println(v)
	}

	ls2 := []string{"g", "a", "m", "a", "n", "i", "a"}
	fmt.Println("yield for loop with generic list param")
	for e := range allLs(ls2) {
		fmt.Println(e)
	}
	for e := range allLs(ls1) {
		fmt.Println(e)
	}

	fmt.Println("Seq is an iterator over sequences of individual values")
	fmt.Println("yield AllSeq")
	for e := range All1(ls1) {
		fmt.Println(e)
	}
	for e := range All1(ls2) {
		fmt.Println(e)
	}

}


func seqYield(yield func(int) bool) {
	for i := 6; i < 11; i++ {
		if !yield(i) {
			break
		}
	}
}

func allInt(n int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

func allLsInt(ns []int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for _, n := range ns {
			if !yield(n) {
				break
			}
		}
	}
}


type Element interface {
	int | string
}

func allLs[E Element](ele []E) func(yield func(E) bool) {
	return func(yield func(E) bool) {
		for _, e := range ele {
			if !yield(e) {
				break
			}
		}
	}
}

func allLsV[V any](vs []V) func(yield func(V) bool) {
	return func(yield func(V) bool) {
		for _, v := range vs {
			if !yield(v) {
				break
			}
		}
	}
}

type Seq[V any] func(yield func(V) bool)

func All1[V any](s []V) Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				break
			}
		}
	}
}