package main

import (
	"fmt"
	"iter"
)

func main() {

	ns := []int{10, 20, 30, 40}
	for v1, v2 := range Pair(All(ns)) {
		fmt.Println(v1, v2)
	}


}

// type Seq0 func(yield func() bool) bool
// type Seq[V any] func(yield func(V) bool) bool
// type Seq2[K, V any] func(yield func(K, V) bool) bool

func All[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				break
			}
		}
	}
}

func Pair[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, _ := next(); v2, ok := next()
			if !ok || !yield(v1, v2) { break }
		}
	}
}