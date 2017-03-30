package main

// #cgo CXXFLAGS: -I$DAALINCLUDE
// #cgo LDFLAGS: -L$DAALLIB -ldaal_core -ldaal_sequential -lpthread -lm
// #include "cholesky.hxx"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {

	// Define the input matrix as an array.
	inputArray := [9]float64{
		1.0, 2.0, 4.0,
		2.0, 13.0, 23.0,
		4.0, 23.0, 77.0,
	}

	// Get the first Cholesky decomposition factor.
	data := (*C.double)(unsafe.Pointer(&inputArray[0]))
	factor := C.choleskyDecompose(3, data)

	// Output the first Cholesky dcomposition factor to stdout.
	fmt.Printf("The first Cholesky decomp. factor is: %d\n", factor)
}
