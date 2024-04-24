using LinearAlgebra
using BenchmarkTools

n = 2000; # 1000
A = randn(n, n);

@btime eigvecs(A)