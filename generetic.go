package main

import (
	"fmt"
)

func SumIntsOrFloats[K comparable, V int64 | float64 | string](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

type Number interface {
    int64 | float64 | string
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}


type Complex interface {
	~complex64 | ~complex128
}

type Float interface {
    ~float32 | ~float64
}

type Unsigned interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func main() {
    // Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	
	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	strings	:= map[string]string{
		"first": "Hello",
		"second": ", ",
		"third": "World",
	}
	
	fmt.Printf("Generic Sums: %v, %v, %v\n",
    	SumIntsOrFloats[string, int64](ints),
    	SumIntsOrFloats[string, float64](floats),
    	SumIntsOrFloats[string, string](strings))

	fmt.Printf("Interface Sums: %v, %v, %v\n",
		SumNumbers[string, int64](ints),
		SumNumbers[string, float64](floats),
		SumNumbers[string, string](strings))
}