package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main() {
	n := 5000

	A := mat.NewDense(n, n, nil)
	b := mat.NewVecDense(n, nil)

	// Fill A and b randomly
	for i := 0; i < n; i++ {
		b.SetVec(i, rand.Float64())
		for j := 0; j < n; j++ {
			A.Set(i, j, rand.Float64())
		}
	}


	start := time.Now()

	var x mat.VecDense
	err := x.SolveVec(A, b)
	if err != nil {
		panic("Failed")
	}

	end := time.Now()

	fmt.Printf("%v", end.Sub(start))
}