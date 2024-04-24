package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main() {
	n := 1000

	A := mat.NewDense(n, n, nil)

	// Fill A randomly
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, rand.Float64())
		}
	}

	start := time.Now()

	ok := new(mat.Eigen).Factorize(A, mat.EigenNone)
	if !ok {
		panic("Failed")
	}
	
	end := time.Now()

	fmt.Printf("%v", end.Sub(start))
}