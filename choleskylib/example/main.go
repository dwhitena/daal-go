package main

import (
	"fmt"

	"github.com/dwhitena/daal-go/choleskylib"
)

func main() {

	// Define the input matrix as an array.
	inputArray := [9]float64{
		1.0, 2.0, 4.0,
		2.0, 13.0, 23.0,
		4.0, 23.0, 77.0,
	}

	// Get the first Cholesky decomposition factor.
	factor := cholesky.CholeskyDecompose(3, &inputArray[0])

	// Output the first Cholesky dcomposition factor to stdout.
	fmt.Printf("The first Cholesky decomp. factor is: %d\n", factor)
}
