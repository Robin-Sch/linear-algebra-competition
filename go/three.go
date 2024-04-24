package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main() {
	m := 1000
	n := 1000

	A := mat.NewDense(m, n, nil)

	// Fill A randomly
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, rand.Float64())
		}
	}

	start := time.Now()

	ok := new(mat.SVD).Factorize(A, mat.SVDThin)
	if !ok {
		panic("Failed")
	}

	end := time.Now()

	fmt.Printf("%v", end.Sub(start))
}