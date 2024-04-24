using LinearAlgebra
using BenchmarkTools

m = 2000; # 1000
n = 2000; # 1000
A = randn(m, n);

@btime svd(A)