package main

import (
	"fmt"
	"iter"
)

func main() {

	ns := []int{10, 20, 30, 40, 50, 60}
	for v1, v2 := range Pair(All(ns)) {
		fmt.Println(v1, v2)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range KVPair(All2(m)) {
		fmt.Println(k, v)
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

func All2[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}
// iter.Pull2 
// func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func())
func KVPair[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	next, stop := iter.Pull2(seq)
	return func(yield func(K, V) bool) {
		defer stop()
		for {
			k, v, ok := next()
			if !ok || !yield(k, v) { break }
		}
	}
}
